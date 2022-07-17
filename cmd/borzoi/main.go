package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/elga-io/borzoi/internal/app/address"
	"github.com/elga-io/borzoi/internal/app/auth"
	"github.com/elga-io/borzoi/internal/app/auth/password"
	"github.com/elga-io/borzoi/internal/app/client"
	"github.com/elga-io/borzoi/internal/app/user"
	"github.com/elga-io/borzoi/internal/pkg/common"
	"github.com/elga-io/borzoi/internal/pkg/config"
	"github.com/elga-io/borzoi/internal/pkg/database"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"github.com/elga-io/borzoi/internal/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//const version = "0.0.0"

//go:embed all:web
var nextFS embed.FS

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Initializing app...")

	cfg := config.NewConfig()
	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	distFS, err := fs.Sub(nextFS, "web")
	if err != nil {
		log.Fatal().Caller().Msg(err.Error())
	}

	// Create database and cache folder in $HOME/.borzoi path.
	folder, err := common.CreateDataFolder(".borzoi")
	if err != nil {
		panic(err)
	}

	db := database.NewSQLDatabase(cfg.SQLType, filepath.Join(folder, "data"))
	cache := database.NewNoSQLDatabase(cfg.NoSQLType, filepath.Join(folder, "cache"))

	if cfg.Migrate {
		err = db.AutoMigrate(
			entity.Identity{},
			entity.User{},
			entity.Phone{},
			entity.Email{},
			entity.Note{},
			entity.Client{})
		if err != nil {
			log.Fatal().Msg(fmt.Sprintf("in indetities migration: %s", err.Error()))
		}
	}

	mw := middleware.Middleware{Cache: cache, Logger: log.Logger}

	router := mux.NewRouter().SkipClean(true)
	router.Use(mw.CorrelationID)

	apiRouter := router.PathPrefix("/api").Subrouter().StrictSlash(true)
	webRouter := router.PathPrefix("/").Subrouter().StrictSlash(true)

	// Log for web path is so verbose
	// So we need active only for debug purpose
	// or use some cache or whatever like that.
	if cfg.Debug {
		router.Use(mw.Log)
	} else {
		apiRouter.Use(mw.Log)
	}

	// Start services.
	// Each service can separated per mini-service.
	// With HTTP, RPC, MQ or another transport.

	// Authentication service.
	// Root level, like /logout or /check endpoint.
	authService := auth.NewService(db, cache)
	authService.HTTPNew(apiRouter)

	// Password auth service.
	authPasswordService := password.NewService(db, cache)
	authPasswordService.HTTPNew(apiRouter)

	// Users service.
	userService := user.NewService(db, cache)
	userService.HTTPNew(apiRouter)

	// Clients service.
	clientService := client.NewService(cfg, db, cache)
	clientService.HTTPNew(apiRouter)

	// Addresses service.
	addressService := address.NewService(db, cache)
	addressService.HTTPNew(apiRouter)

	// Web app.
	webHandler := http.FileServer(http.FS(distFS))
	webRouter.PathPrefix("").Handler(webHandler)

	// Print routes when startup app.
	if cfg.LogRoutes {
		common.PrintRoutes([]*mux.Router{webRouter, apiRouter})
	}

	srv := &http.Server{
		Addr:         ":" + cfg.HTTPPort,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Info().Msg(fmt.Sprintf("started server on :%[1]s, url: http://localhost:%[1]s", cfg.HTTPPort))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Msg(fmt.Sprintf("listen: %s", err.Error()))
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Info().Msg("shutting down gracefully. Press ctrl+c again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Msg(fmt.Sprintf("server forced to shutdown: %s", err.Error()))
	}
	log.Info().Msg("server exiting")
}

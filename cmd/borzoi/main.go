package main

import (
	"context"
	"embed"
	"errors"
	"flag"
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
	"github.com/elga-io/borzoi/internal/pkg/config"
	"github.com/elga-io/borzoi/internal/pkg/database"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"github.com/elga-io/borzoi/internal/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//const version = "0.0.0"

//go:embed web
//go:embed web/_next/static
//go:embed web/_next/static/chunks/pages/*.js
//go:embed web/_next/static/*/*.js
var nextFS embed.FS

func main() {
	fMigrate := flag.Bool("migrate", false, "Enable GORM migration")

	flag.Parse()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Initializing app...")

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	distFS, err := fs.Sub(nextFS, "web")
	if err != nil {
		log.Fatal().Caller().Msg(err.Error())
	}

	// Create database and cache folder in $HOME/.borzoi path.
	folder, err := createDataFolder(".borzoi")
	if err != nil {
		panic(err)
	}

	cfg := config.NewConfig()
	db := database.NewSQLDatabase("sqlite", filepath.Join(folder, "data"))
	cache := database.NewNoSQLDatabase("badger", filepath.Join(folder, "cache"))

	if *fMigrate {
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
	router.Use(mw.Log)

	apiRouter := router.PathPrefix("/api").Subrouter().StrictSlash(true)
	webRouter := router.PathPrefix("/").Subrouter().StrictSlash(true)

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
	clientService := client.NewService(db, cache)
	clientService.HTTPNew(apiRouter)

	// Addresses service.
	addressService := address.NewService(db, cache)
	addressService.HTTPNew(apiRouter)

	// Web app.
	webHandler := http.FileServer(http.FS(distFS))
	webRouter.PathPrefix("").Handler(webHandler)

	// Print routes when startup app.
	printRoutes([]*mux.Router{webRouter, apiRouter})

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

func createDataFolder(name string) (folder string, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	folder = filepath.Join(home, name)
	if _, err = os.Stat(folder); os.IsNotExist(err) {
		err = os.Mkdir(folder, os.ModePerm)
		if err != nil {
			return
		}
	} else if err != nil {
		return
	}
	return
}

func printRoutes(rs []*mux.Router) {
	for _, r := range rs {
		_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			uri, err := route.GetPathTemplate()
			if err != nil {
				log.Error().Msg(fmt.Sprintf("with get path template: %s", err.Error()))
				return err
			}

			method, err := route.GetMethods()
			if err != nil {
				if errors.Is(err, mux.ErrMethodMismatch) {
					return err
				}
			}

			if uri != "" && len(method) != 0 {
				log.Info().Msg(fmt.Sprintf("%s %s", uri, method))
			}
			return nil
		})
	}
}

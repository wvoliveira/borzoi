package main

import (
	"context"
	"embed"
	"errors"
	"flag"
	"fmt"
	"github.com/elga-io/borzoi/internal/app/auth"
	"github.com/elga-io/borzoi/internal/app/auth/password"
	"github.com/elga-io/borzoi/internal/app/user"
	"github.com/elga-io/borzoi/internal/pkg/config"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"github.com/elga-io/borzoi/internal/pkg/middleware"
	"github.com/elga-io/canideos/database"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

//const version = "0.0.0"

////go:embed web
////go:embed web/_next/static
////go:embed web/_next/static/chunks/pages/*.js
////go:embed web/_next/static/*/*.js
var nextFS embed.FS

func main() {
	flag.Parse()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Initializing app...")

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	distFS, err := fs.Sub(nextFS, "web")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	cfg := config.NewConfig()
	db := database.NewSQLDatabase("sqlite", "./.db/data")
	cache := database.NewNoSQLDatabase("badger", "./.db/cache")

	if *fMigrate {
		err = db.AutoMigrate(entity.Identity{}, entity.User{})
		if err != nil {
			log.Fatal().Msg(fmt.Sprintf("in indetities migration: %s", err.Error()))
		}
	}

	mw := middleware.Middleware{Cache: cache, Logger: log.Logger}

	router := mux.NewRouter().SkipClean(true)
	router.Use(mw.CorrelationID)
	router.Use(mw.Log)

	webRouter := router.MatcherFunc(func(req *http.Request, match *mux.RouteMatch) bool {
		return req.RequestURI == "/" ||
			strings.HasPrefix(req.RequestURI, "/_next") ||
			req.RequestURI == "/favicon.ico"
	}).Subrouter().StrictSlash(true)

	apiRouter := router.PathPrefix("/api").Subrouter().StrictSlash(true)

	authService := auth.NewService(db, cache)
	authService.HTTPNew(apiRouter)

	authPasswordService := password.NewService(db, cache)
	authPasswordService.HTTPNew(apiRouter)

	userService := user.NewService(db, cache)
	userService.HTTPNew(apiRouter)

	webHandler := http.FileServer(http.FS(distFS))
	webRouter.PathPrefix("").Handler(webHandler)

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

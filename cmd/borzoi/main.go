package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"github.com/elga-io/borzoi/internal/app/auth"
	"github.com/elga-io/borzoi/internal/app/auth/password"
	"github.com/elga-io/borzoi/internal/app/user"
	"github.com/elga-io/borzoi/internal/pkg/config"
	"github.com/elga-io/borzoi/internal/pkg/entity"
	"github.com/elga-io/canideos/database"
	"github.com/gorilla/mux"
	"io/fs"
	"log"
	"net/http"
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

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	distFS, err := fs.Sub(nextFS, "web")
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.NewConfig()
	db := database.NewSQLDatabase("sqlite", "./.db/data")
	cache := database.NewNoSQLDatabase("badger", "./.db/cache")

	if *fMigrate {
		err = db.AutoMigrate(entity.Identity{}, entity.User{})
		if err != nil {
			log.Fatal(fmt.Sprintf("%s error: in indetities migration: %s\n", time.Now(), err.Error()))
		}
	}

	router := mux.NewRouter().SkipClean(true)
	router.Use(middlewareLog)

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

	printRoutes([]*mux.Router{apiRouter})

	srv := &http.Server{
		Addr:         ":" + cfg.HTTPPort,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		fmt.Printf("started server on :%[1]s, url: http://localhost:%[1]s\n", cfg.HTTPPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n\n", err.Error())
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %s\n\n", err.Error())
	}
	fmt.Println("Server exiting")
}

func middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: get status code from http.ResponseWriter
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func printRoutes(rs []*mux.Router) {
	for _, r := range rs {
		_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			t, err := route.GetPathTemplate()
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			if t != "" {
				fmt.Println(t)
			}
			return nil
		})
	}
}

package common

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

// CreateDataFolder create app folder in user home.
// Example: /home/user/.borzoi/
func CreateDataFolder(name string) (folder string, err error) {
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

// PrintRoutes just print path and methods allowed by routes.
func PrintRoutes(rs []*mux.Router) {
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

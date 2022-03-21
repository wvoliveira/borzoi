package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func New() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

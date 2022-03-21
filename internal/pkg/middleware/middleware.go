package middleware

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/elga-io/borzoi/internal/pkg/constant"
	"github.com/elga-io/borzoi/internal/pkg/unique"
	"github.com/elga-io/borzoi/internal/pkg/user"
	e "github.com/elga-io/canideos/errors"
	"github.com/google/uuid"
	zl "github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type LogResponseWriter struct {
	http.ResponseWriter
	statusCode int
	buf        bytes.Buffer
}

func NewLogResponseWriter(w http.ResponseWriter) *LogResponseWriter {
	return &LogResponseWriter{ResponseWriter: w}
}

func (w *LogResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *LogResponseWriter) Write(body []byte) (int, error) {
	w.buf.Write(body)
	return w.ResponseWriter.Write(body)
}

type Middleware struct {
	Cache *badger.DB
}

// Auth checks if the client is authenticated
func (m Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := unique.GetUUID(r.Context())
		cookie, err := r.Cookie("session")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				zl.Warn().Caller().Str("id", id).Str("text", http.ErrNoCookie.Error()).Msg("service")
				e.EncodeError(w, e.ErrUnauthorized)
				return
			}
			zl.Error().Caller().Msg(fmt.Sprintf("to get session from cookie: %s", err.Error()))
			e.EncodeError(w, err)
			return
		}

		key := fmt.Sprintf(constant.KeySession, cookie.Value)
		u, err := user.GetUserBySession(m.Cache, []byte(key))
		if err != nil {
			e.EncodeError(w, err)
			return
		}

		if u.ID == "" {
			e.EncodeError(w, e.ErrUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UUID add unique ID for each request.
func UUID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Log print request info to stdout.
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := unique.GetUUID(ctx)
		startTime := time.Now()
		logRespWriter := NewLogResponseWriter(w)

		zl.
			Info().
			Caller().
			Str("id", id).
			Str("remote_addr", r.RemoteAddr).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Msg("request")

		next.ServeHTTP(logRespWriter, r)

		zl.
			Info().
			Caller().
			Str("id", id).
			Str("duration", time.Since(startTime).String()).
			Int("status", logRespWriter.statusCode).
			Msg("response")
	})
}

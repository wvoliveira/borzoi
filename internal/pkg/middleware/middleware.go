package middleware

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/elga-io/borzoi/internal/pkg/session"
	e "github.com/elga-io/canideos/errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
	Cache  *badger.DB
	Logger zerolog.Logger
}

// Auth checks if the client is authenticated
func (m Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := log.Ctx(r.Context())

		cookie, err := r.Cookie("session")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				l.Warn().Caller().Msg(http.ErrNoCookie.Error())
				e.EncodeError(w, e.ErrAuthUnauthorized)
				return
			}
			l.Error().Caller().Msg(fmt.Sprintf("to get session from cookie: %s", err.Error()))
			e.EncodeError(w, err)
			return
		}

		key := fmt.Sprintf(session.CacheKey, cookie.Value)
		userID, err := session.UserGetIDFromSession(r.Context(), m.Cache, []byte(key))
		if err != nil {
			l.Error().Caller().Msg(err.Error())
			e.EncodeError(w, err)
			return
		}

		if userID == "" {
			e.EncodeError(w, e.ErrAuthUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// CorrelationID add unique ID for each request.
func (m Middleware) CorrelationID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := r.Header.Get("X-Correlation-Id")
		if id == "" {
			id = uuid.New().String()
		}

		ctx = context.WithValue(ctx, "correlation_id", id)
		r = r.WithContext(ctx)

		m.Logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("correlation_id", id)
		})
		r = r.WithContext(m.Logger.WithContext(r.Context()))

		w.Header().Set("X-Correlation-Id", id)
		next.ServeHTTP(w, r)
	})
}

// Log print request info to stdout.
func (m Middleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		logRespWriter := NewLogResponseWriter(w)

		l := log.Ctx(r.Context())
		l.
			Info().
			Caller().
			Str("remote_addr", r.RemoteAddr).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Msg("request")

		next.ServeHTTP(logRespWriter, r)

		l.
			Info().
			Caller().
			Str("duration", time.Since(startTime).String()).
			Int("status", logRespWriter.statusCode).
			Msg("response")
	})
}

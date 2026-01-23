package logy

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reqID := middleware.GetReqID(r.Context())
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
		Global.Info().
			Str("id", reqID).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("status", ww.Status()).
			Dur("latency", time.Since(start)).
			Str("ip", r.RemoteAddr).
			Msg("http_request")
	})
}

func Middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		Global.Info().
			Str("method", r.Method).
			Str("uri", r.RequestURI).
			Str("ip", r.RemoteAddr).
			Str("id", middleware.GetReqID(r.Context())).
			Int("status", ww.Status()).
			Int("bytes", ww.BytesWritten()).
			Dur("lat", time.Since(start)). // Latency
			Msg("http")
	})
}

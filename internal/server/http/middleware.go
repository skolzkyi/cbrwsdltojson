package internalhttp

import (
	"net/http"
	"strings"
	"time"

	zap "go.uber.org/zap"
)

func (s *Server) loggingMiddleware(next http.HandlerFunc, log Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next.ServeHTTP(w, r)

		pathSl := strings.Split(r.URL.Path, "/")
		s.observeRequestDuration(w.Header().Get("Status"), pathSl[len(pathSl)-1], time.Since(t))

		log.GetZapLogger().With(
			zap.String("Client IP", r.RemoteAddr),
			zap.String("Request DateTime", time.Now().String()),
			zap.String("Method", r.Method),
			zap.String("Request URL", r.RequestURI),
			zap.String("Request Scheme", r.URL.Scheme),
			zap.String("Request Status", w.Header().Get("Status")),
			zap.String("Time of request work", time.Since(t).String()),
			zap.String("Request User-Agent", r.Header.Get("User-Agent")),
		).Info("http middleware log")

		errHeader := w.Header().Get("ErrCustom")
		if errHeader != "" {
			log.Error("Error middleware logging: " + errHeader)
		}
	}
}

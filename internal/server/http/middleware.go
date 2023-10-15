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

		timeDelta := time.Since(t)
		status := w.Header().Get("Status")
		pathSl := strings.Split(r.URL.Path, "/")
		handler := pathSl[len(pathSl)-1]
		s.observeRequestCounterTotal(status, handler)
		s.observeRequestDuration(status, handler, timeDelta)

		log.GetZapLogger().With(
			zap.String("Client IP", r.RemoteAddr),
			zap.String("Request DateTime", time.Now().String()),
			zap.String("Method", r.Method),
			zap.String("Request URL", r.RequestURI),
			zap.String("Request Scheme", r.URL.Scheme),
			zap.String("Request Status", status),
			zap.String("Time of request work", timeDelta.String()),
			zap.String("Request User-Agent", r.Header.Get("User-Agent")),
		).Info("http middleware log")

		errHeader := w.Header().Get("ErrCustom")
		if errHeader != "" {
			log.Error("Error middleware logging: " + errHeader)
		}
	}
}

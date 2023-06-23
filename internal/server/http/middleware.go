package internalhttp

import (
	"net/http"
	"time"

	zap "go.uber.org/zap"
)


func loggingMiddleware(next http.HandlerFunc, log Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next.ServeHTTP(w, r)
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
		
		errHeader:= w.Header().Get("ErrCustom")
		if errHeader != "" {
			log.Error("Error middleware logging: "+errHeader)
		}
	}
}

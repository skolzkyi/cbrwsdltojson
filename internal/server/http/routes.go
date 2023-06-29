package internalhttp

import (
	"net/http"
)

func (s *Server) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/GetMethodDataWithoutCache/", loggingMiddleware(s.GetMethodDataWithoutCache, s.logg))
	mux.HandleFunc("/GetCursOnDate", loggingMiddleware(s.GetCursOnDate, s.logg))

	return mux
}

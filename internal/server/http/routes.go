package internalhttp

import (
	"net/http"
)

func (s *Server) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/GetCursOnDate", loggingMiddleware(s.GetCursOnDate, s.logg))

	return mux
}

package internalhttp

import (
	"net/http"
)

func (s *Server) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/GetMethodDataWithoutCache/", loggingMiddleware(s.GetMethodDataWithoutCache, s.logg))

	mux.HandleFunc("/GetCursOnDateXML", loggingMiddleware(s.GetCursOnDateXML, s.logg))
	mux.HandleFunc("/BiCurBaseXML", loggingMiddleware(s.BiCurBaseXML, s.logg))

	return mux
}

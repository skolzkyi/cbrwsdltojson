package internalhttp

import (
	"net/http"
)

func (s *Server) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/GetMethodDataWithoutCache/", loggingMiddleware(s.GetMethodDataWithoutCache, s.logg))

	mux.HandleFunc("/GetCursOnDateXML", loggingMiddleware(s.GetCursOnDateXML, s.logg))
	mux.HandleFunc("/BiCurBaseXML", loggingMiddleware(s.BiCurBaseXML, s.logg))
	mux.HandleFunc("/BliquidityXML", loggingMiddleware(s.BliquidityXML, s.logg))
	mux.HandleFunc("/DepoDynamicXML", loggingMiddleware(s.DepoDynamicXML, s.logg))
	mux.HandleFunc("/DragMetDynamicXML", loggingMiddleware(s.DragMetDynamicXML, s.logg))

	return mux
}

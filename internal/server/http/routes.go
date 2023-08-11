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
	mux.HandleFunc("/DVXML", loggingMiddleware(s.DVXML, s.logg))
	mux.HandleFunc("/EnumReutersValutesXML", loggingMiddleware(s.EnumReutersValutesXML, s.logg))
	mux.HandleFunc("/EnumValutesXML", loggingMiddleware(s.EnumValutesXML, s.logg))
	mux.HandleFunc("/KeyRateXML", loggingMiddleware(s.KeyRateXML, s.logg))
	mux.HandleFunc("/MainInfoXML", loggingMiddleware(s.MainInfoXML, s.logg))
	mux.HandleFunc("/mrrf7DXML", loggingMiddleware(s.Mrrf7DXML, s.logg))
	mux.HandleFunc("/mrrfXML", loggingMiddleware(s.MrrfXML, s.logg))
	mux.HandleFunc("/NewsInfoXML", loggingMiddleware(s.NewsInfoXML, s.logg))

	return mux
}

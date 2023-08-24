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
	mux.HandleFunc("/OmodInfoXML", loggingMiddleware(s.OmodInfoXML, s.logg))
	mux.HandleFunc("/OstatDepoNewXML", loggingMiddleware(s.OstatDepoNewXML, s.logg))
	mux.HandleFunc("/OstatDepoXML", loggingMiddleware(s.OstatDepoXML, s.logg))
	mux.HandleFunc("/OstatDynamicXML", loggingMiddleware(s.OstatDynamicXML, s.logg))
	mux.HandleFunc("/OvernightXML", loggingMiddleware(s.OvernightXML, s.logg))
	mux.HandleFunc("/RepoDebtXML", loggingMiddleware(s.RepoDebtXML, s.logg))
	mux.HandleFunc("/RepoDebtUSDXML", loggingMiddleware(s.RepoDebtUSDXML, s.logg))
	mux.HandleFunc("/ROISfixXML", loggingMiddleware(s.ROISfixXML, s.logg))
	mux.HandleFunc("/RuoniaSVXML", loggingMiddleware(s.RuoniaSVXML, s.logg))
	mux.HandleFunc("/RuoniaXML", loggingMiddleware(s.RuoniaXML, s.logg))
	mux.HandleFunc("/SaldoXML", loggingMiddleware(s.SaldoXML, s.logg))
	mux.HandleFunc("/SwapDayTotalXML", loggingMiddleware(s.SwapDayTotalXML, s.logg))
	mux.HandleFunc("/SwapDynamicXML", loggingMiddleware(s.SwapDynamicXML, s.logg))
	mux.HandleFunc("/SwapInfoSellUSDVolXML", loggingMiddleware(s.SwapInfoSellUSDVolXML, s.logg))
	mux.HandleFunc("/SwapInfoSellUSDXML", loggingMiddleware(s.SwapInfoSellUSDXML, s.logg))
	mux.HandleFunc("/SwapInfoSellVolXML", loggingMiddleware(s.SwapInfoSellVolXML, s.logg))
	mux.HandleFunc("/SwapInfoSellXML", loggingMiddleware(s.SwapInfoSellXML, s.logg))

	return mux
}

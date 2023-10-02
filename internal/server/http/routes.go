package internalhttp

import (
	"net/http"
)

func (s *Server) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/GetMethodDataWithoutCache/", s.loggingMiddleware(s.GetMethodDataWithoutCache, s.logg))

	mux.HandleFunc("/AllDataInfoXML", s.loggingMiddleware(s.AllDataInfoXML, s.logg))
	mux.HandleFunc("/GetCursOnDateXML", s.loggingMiddleware(s.GetCursOnDateXML, s.logg))
	mux.HandleFunc("/BiCurBaseXML", s.loggingMiddleware(s.BiCurBaseXML, s.logg))
	mux.HandleFunc("/BliquidityXML", s.loggingMiddleware(s.BliquidityXML, s.logg))
	mux.HandleFunc("/DepoDynamicXML", s.loggingMiddleware(s.DepoDynamicXML, s.logg))
	mux.HandleFunc("/DragMetDynamicXML", s.loggingMiddleware(s.DragMetDynamicXML, s.logg))
	mux.HandleFunc("/DVXML", s.loggingMiddleware(s.DVXML, s.logg))
	mux.HandleFunc("/EnumReutersValutesXML", s.loggingMiddleware(s.EnumReutersValutesXML, s.logg))
	mux.HandleFunc("/EnumValutesXML", s.loggingMiddleware(s.EnumValutesXML, s.logg))
	mux.HandleFunc("/KeyRateXML", s.loggingMiddleware(s.KeyRateXML, s.logg))
	mux.HandleFunc("/MainInfoXML", s.loggingMiddleware(s.MainInfoXML, s.logg))
	mux.HandleFunc("/mrrf7DXML", s.loggingMiddleware(s.Mrrf7DXML, s.logg))
	mux.HandleFunc("/mrrfXML", s.loggingMiddleware(s.MrrfXML, s.logg))
	mux.HandleFunc("/NewsInfoXML", s.loggingMiddleware(s.NewsInfoXML, s.logg))
	mux.HandleFunc("/OmodInfoXML", s.loggingMiddleware(s.OmodInfoXML, s.logg))
	mux.HandleFunc("/OstatDepoNewXML", s.loggingMiddleware(s.OstatDepoNewXML, s.logg))
	mux.HandleFunc("/OstatDepoXML", s.loggingMiddleware(s.OstatDepoXML, s.logg))
	mux.HandleFunc("/OstatDynamicXML", s.loggingMiddleware(s.OstatDynamicXML, s.logg))
	mux.HandleFunc("/OvernightXML", s.loggingMiddleware(s.OvernightXML, s.logg))
	mux.HandleFunc("/RepoDebtXML", s.loggingMiddleware(s.RepoDebtXML, s.logg))
	mux.HandleFunc("/RepoDebtUSDXML", s.loggingMiddleware(s.RepoDebtUSDXML, s.logg))
	mux.HandleFunc("/ROISfixXML", s.loggingMiddleware(s.ROISfixXML, s.logg))
	mux.HandleFunc("/RuoniaSVXML", s.loggingMiddleware(s.RuoniaSVXML, s.logg))
	mux.HandleFunc("/RuoniaXML", s.loggingMiddleware(s.RuoniaXML, s.logg))
	mux.HandleFunc("/SaldoXML", s.loggingMiddleware(s.SaldoXML, s.logg))
	mux.HandleFunc("/SwapDayTotalXML", s.loggingMiddleware(s.SwapDayTotalXML, s.logg))
	mux.HandleFunc("/SwapDynamicXML", s.loggingMiddleware(s.SwapDynamicXML, s.logg))
	mux.HandleFunc("/SwapInfoSellUSDVolXML", s.loggingMiddleware(s.SwapInfoSellUSDVolXML, s.logg))
	mux.HandleFunc("/SwapInfoSellUSDXML", s.loggingMiddleware(s.SwapInfoSellUSDXML, s.logg))
	mux.HandleFunc("/SwapInfoSellVolXML", s.loggingMiddleware(s.SwapInfoSellVolXML, s.logg))
	mux.HandleFunc("/SwapInfoSellXML", s.loggingMiddleware(s.SwapInfoSellXML, s.logg))
	mux.HandleFunc("/SwapMonthTotalXML", s.loggingMiddleware(s.SwapMonthTotalXML, s.logg))

	return mux
}

package internalhttp

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"

	helpers "github.com/skolzkyi/cbrwsdltojson/helpers"
	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
)

type requestData interface {
	Init()
	Validate() error
}

type blLayerMethod func(context.Context, interface{}, string) (interface{}, error)

// without parameters.
type blLayerMethodWP func(context.Context) (interface{}, error)

var (
	ErrInJSONBadParse        = errors.New("error parsing input json")
	ErrOutJSONBadParse       = errors.New("error parsing output json")
	ErrUnsupportedMethod     = errors.New("http unsupported method")
	ErrNoSOAPActionInRequest = errors.New("no SOAPAction in request")
)

func apiErrHandler(err error, w *http.ResponseWriter) {
	var errMessage string
	if err != nil {
		W := *w
		if errors.Is(err, datastructures.ErrBadInputDateData) || errors.Is(err, datastructures.ErrBadRawData) {
			errMessage = helpers.StringBuild(http.StatusText(http.StatusBadRequest), " (", err.Error(), ")")
			http.Error(W, errMessage, http.StatusBadRequest)
			W.Header().Add("Status", "400")
		} else {
			errMessage = helpers.StringBuild(http.StatusText(http.StatusInternalServerError), " (", err.Error(), ")")
			http.Error(W, errMessage, http.StatusInternalServerError)
			W.Header().Add("Status", "500")
		}
		W.Header().Add("ErrCustom", err.Error())
	}
}

func (s *Server) GetMethodDataWithoutCache(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	switch r.Method {
	case http.MethodPost:
		path := strings.Trim(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 2 {
			apiErrHandler(ErrNoSOAPActionInRequest, &w)
			return
		}

		SOAPAction := pathParts[1]

		body, err := io.ReadAll(r.Body)
		if err != nil {
			apiErrHandler(err, &w)
			return
		}
		rawBody := helpers.ClearStringByWhitespaceAndLinebreak(string(body))
		s.app.RemoveDataInMemCacheBySOAPAction(SOAPAction + rawBody)

		// 307, not 303: on 307 no lost body and no change verb to GET
		http.Redirect(w, r, "/"+SOAPAction, http.StatusTemporaryRedirect)

	default:
		apiErrHandler(ErrUnsupportedMethod, &w)
		return
	}
}

func (s *Server) universalMethodHandler(w http.ResponseWriter, r *http.Request, reqData requestData, appMethod blLayerMethod) {
	defer r.Body.Close()
	fullRequestTimeout, err := s.GetFullRequestTimeout()
	if err != nil {
		apiErrHandler(err, &w)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), fullRequestTimeout)
	defer cancel()
	switch r.Method {
	case http.MethodPost:
		body, err := s.ReadDataFromInputJSON(reqData, r)
		if err != nil {
			apiErrHandler(err, &w)
			return
		}

		reqData.Init()

		err = reqData.Validate()
		if err != nil {
			apiErrHandler(err, &w)
			return
		}

		answer, err := appMethod(ctx, reqData, body)
		if err != nil {
			apiErrHandler(err, &w)
			return
		}

		err = s.WriteDataToOutputJSON(answer, w)
		if err != nil {
			apiErrHandler(err, &w)
			return
		}
		w.Header().Add("Status", "200")
		return

	default:
		apiErrHandler(ErrUnsupportedMethod, &w)
		return
	}
}

// without parameters.
func (s *Server) universalMethodHandlerWP(w http.ResponseWriter, r *http.Request, appMethod blLayerMethodWP) {
	defer r.Body.Close()
	fullRequestTimeout, err := s.GetFullRequestTimeout()
	if err != nil {
		apiErrHandler(err, &w)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), fullRequestTimeout)
	defer cancel()
	switch r.Method {
	case http.MethodPost:

		answer, err := appMethod(ctx)
		if err != nil {
			apiErrHandler(err, &w)
			return
		}

		err = s.WriteDataToOutputJSON(answer, w)
		if err != nil {
			apiErrHandler(err, &w)
			return
		}
		w.Header().Add("Status", "200")
		return

	default:
		apiErrHandler(ErrUnsupportedMethod, &w)
		return
	}
}

func (s *Server) AllDataInfoXML(w http.ResponseWriter, r *http.Request) {
	s.universalMethodHandlerWP(w, r, s.app.AllDataInfoXML)
}

func (s *Server) GetCursOnDateXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.GetCursOnDateXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.GetCursOnDateXML)
}

func (s *Server) BiCurBaseXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.BiCurBaseXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.BiCurBaseXML)
}

func (s *Server) BliquidityXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.BliquidityXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.BliquidityXML)
}

func (s *Server) DepoDynamicXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.DepoDynamicXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.DepoDynamicXML)
}

func (s *Server) DragMetDynamicXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.DragMetDynamicXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.DragMetDynamicXML)
}

func (s *Server) DVXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.DVXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.DVXML)
}

func (s *Server) EnumReutersValutesXML(w http.ResponseWriter, r *http.Request) {
	s.universalMethodHandlerWP(w, r, s.app.EnumReutersValutesXML)
}

func (s *Server) EnumValutesXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.EnumValutesXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.EnumValutesXML)
}

func (s *Server) KeyRateXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.KeyRateXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.KeyRateXML)
}

func (s *Server) MainInfoXML(w http.ResponseWriter, r *http.Request) {
	s.universalMethodHandlerWP(w, r, s.app.MainInfoXML)
}

func (s *Server) Mrrf7DXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.Mrrf7DXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.Mrrf7DXML)
}

func (s *Server) MrrfXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.MrrfXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.MrrfXML)
}

func (s *Server) NewsInfoXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.NewsInfoXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.NewsInfoXML)
}

func (s *Server) OmodInfoXML(w http.ResponseWriter, r *http.Request) {
	s.universalMethodHandlerWP(w, r, s.app.OmodInfoXML)
}

func (s *Server) OstatDepoNewXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.OstatDepoNewXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.OstatDepoNewXML)
}

func (s *Server) OstatDepoXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.OstatDepoXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.OstatDepoXML)
}

func (s *Server) OstatDynamicXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.OstatDynamicXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.OstatDynamicXML)
}

func (s *Server) OvernightXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.OvernightXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.OvernightXML)
}

func (s *Server) RepoDebtXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.Repo_debtXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.RepoDebtXML)
}

func (s *Server) RepoDebtUSDXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.RepoDebtUSDXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.RepoDebtUSDXML)
}

func (s *Server) ROISfixXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.ROISfixXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.ROISfixXML)
}

func (s *Server) RuoniaSVXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.RuoniaSVXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.RuoniaSVXML)
}

func (s *Server) RuoniaXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.RuoniaXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.RuoniaXML)
}

func (s *Server) SaldoXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.SaldoXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.SaldoXML)
}

func (s *Server) SwapDayTotalXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.SwapDayTotalXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.SwapDayTotalXML)
}

func (s *Server) SwapDynamicXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.SwapDynamicXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.SwapDynamicXML)
}

func (s *Server) SwapInfoSellUSDVolXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.SwapInfoSellUSDVolXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.SwapInfoSellUSDVolXML)
}

func (s *Server) SwapInfoSellUSDXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.SwapInfoSellUSDXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.SwapInfoSellUSDXML)
}

func (s *Server) SwapInfoSellVolXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.SwapInfoSellVolXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.SwapInfoSellVolXML)
}

func (s *Server) SwapInfoSellXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.SwapInfoSellXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.SwapInfoSellXML)
}

func (s *Server) SwapMonthTotalXML(w http.ResponseWriter, r *http.Request) {
	newRequest := datastructures.SwapMonthTotalXML{}
	s.universalMethodHandler(w, r, &newRequest, s.app.SwapMonthTotalXML)
}

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

var (
	ErrInJSONBadParse        = errors.New("error parsing input json")
	ErrOutJSONBadParse       = errors.New("error parsing output json")
	ErrUnsupportedMethod     = errors.New("http unsupported method")
	ErrNoSOAPActionInRequest = errors.New("no SOAPAction in request")
)

func apiErrHandler(err error, w *http.ResponseWriter) {
	if err != nil {
		W := *w
		errMessage := helpers.StringBuild(http.StatusText(http.StatusInternalServerError), " (", err.Error(), ")")
		http.Error(W, errMessage, http.StatusInternalServerError)
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
		return

	default:
		apiErrHandler(ErrUnsupportedMethod, &w)
		return
	}
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

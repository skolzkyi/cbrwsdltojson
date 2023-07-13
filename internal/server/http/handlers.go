package internalhttp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	helpers "github.com/skolzkyi/cbrwsdltojson/helpers"
	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
)

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

		fmt.Println("GetMethodDataWithoutCache SOAPAction: ", SOAPAction)

		s.app.RemoveDataInMemCacheBySOAPAction(SOAPAction)

		// 307, not 303: on 307 not lost body and not change verb to GET
		http.Redirect(w, r, "/"+SOAPAction, http.StatusTemporaryRedirect)

	default:
		apiErrHandler(ErrUnsupportedMethod, &w)
		return
	}
}

func (s *Server) GetCursOnDateXML(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx, cancel := context.WithTimeout(r.Context(), s.Config.GetCBRWSDLTimeout())
	defer cancel()
	switch r.Method {
	case http.MethodPost:
		newRequest := datastructures.GetCursOnDateXML{}

		err := s.ReadDataFromInputJSON(&newRequest, r)
		if err != nil {
			apiErrHandler(err, &w)
			return
		}
		err = newRequest.Validate(s.Config.GetDateTimeRequestLayout())
		if err != nil {
			apiErrHandler(err, &w)
			return
		}

		answer, err := s.app.GetCursOnDateXML(ctx, newRequest)
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

func (s *Server) BiCurBaseXML(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx, cancel := context.WithTimeout(r.Context(), s.Config.GetCBRWSDLTimeout())
	defer cancel()
	switch r.Method {
	case http.MethodPost:
		newRequest := datastructures.BiCurBaseXML{}

		err := s.ReadDataFromInputJSON(&newRequest, r)
		if err != nil {
			apiErrHandler(err, &w)
			return
		}
		err = newRequest.Validate(s.Config.GetDateTimeRequestLayout())
		if err != nil {
			apiErrHandler(err, &w)
			return
		}

		answer, err := s.app.BiCurBaseXML(ctx, newRequest)
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

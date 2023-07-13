package internalhttp

import (
	"context"
	"errors"
	"fmt"
	"io"
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

		body, err := io.ReadAll(r.Body)
		if err != nil {
			apiErrHandler(err, &w)
			return
		}
		rawBody := helpers.ClearStringByWhitespaceAndLinebreak(string(body))
		fmt.Println("nocache_tag_in_upd_add: ", SOAPAction+rawBody)
		s.app.RemoveDataInMemCacheBySOAPAction(SOAPAction + rawBody)

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

		body, err := s.ReadDataFromInputJSON(&newRequest, r)
		if err != nil {
			apiErrHandler(err, &w)
			return
		}
		err = newRequest.Validate(s.Config.GetDateTimeRequestLayout())
		if err != nil {
			apiErrHandler(err, &w)
			return
		}

		answer, err := s.app.GetCursOnDateXML(ctx, newRequest, body)
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

		body, err := s.ReadDataFromInputJSON(&newRequest, r)
		if err != nil {
			apiErrHandler(err, &w)
			return
		}
		err = newRequest.Validate(s.Config.GetDateTimeRequestLayout())
		if err != nil {
			apiErrHandler(err, &w)
			return
		}

		answer, err := s.app.BiCurBaseXML(ctx, newRequest, body)
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

package internalhttp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	//"strconv"

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
	W := *w
	if err != nil {
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

func (s *Server) GetCursOnDate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx, cancel := context.WithTimeout(r.Context(), s.Config.GetCBRWSDLTimeout())
	defer cancel()
	fmt.Println("GetCursOnDate method: ", r.Method)
	switch r.Method {
	case http.MethodPost:
		newRequest := datastructures.GetCursOnDateXML{}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("err io.ReadAll")
			apiErrHandler(err, &w)
			return
		}
		fmt.Println("io.ReadAll: ", string(body))
		err = json.Unmarshal(body, &newRequest)
		if err != nil {
			fmt.Println("err json.Unmarshal")
			apiErrHandler(err, &w)
			return
		}

		fmt.Println("newRequest: ", newRequest)

		err = newRequest.Validate(s.Config.GetDateTimeRequestLayout())
		if err != nil {
			fmt.Println("newRequest.Validate(): ", err.Error())
			apiErrHandler(err, &w)
			return
		}

		err, answer := s.app.GetCursOnDate(ctx, newRequest)
		if err != nil {
			apiErrHandler(err, &w)
			return
		}

		jsonstring, err := json.Marshal(answer)
		if err != nil {
			apiErrHandler(err, &w)
			return
		}
		_, err = w.Write(jsonstring)
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

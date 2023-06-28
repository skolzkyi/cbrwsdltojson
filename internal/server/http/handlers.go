package internalhttp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	//"strconv"

	helpers "github.com/skolzkyi/cbrwsdltojson/helpers"
	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
)

var (
	ErrInJSONBadParse    = errors.New("error parsing input json")
	ErrOutJSONBadParse   = errors.New("error parsing output json")
	ErrUnsupportedMethod = errors.New("http unsupported method")
)

func apiErrHandler(err error, w *http.ResponseWriter) {
	W := *w
	if err != nil {
		errMessage := helpers.StringBuild(http.StatusText(http.StatusInternalServerError), " (", err.Error(), ")")
		http.Error(W, errMessage, http.StatusInternalServerError)
	}
}

func (s *Server) GetCursOnDate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx, cancel := context.WithTimeout(r.Context(), s.Config.GetCBRWSDLTimeout())
	defer cancel()

	switch r.Method {
	case http.MethodPost:
		newRequest := datastructures.GetCursOnDateXML{}
		jsonstring1, err := json.Marshal(newRequest)
		fmt.Println("jsonstring1: ", string(jsonstring1))

		body, err := io.ReadAll(r.Body)
		if err != nil {
			apiErrHandler(err, &w)
			return
		}

		err = json.Unmarshal(body, &newRequest)
		if err != nil {
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

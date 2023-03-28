//Filename cmd/api/helpers.go

package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]interface{}

func (app *application) readCodeParams(r *http.Request) (string, error) {
	params := httprouter.ParamsFromContext(r.Context())
	code := params.ByName("id")

	return code, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	//convert to a JSON format
	js, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return err
	}

	js = append(js, '\n')

	//Add any headers sent
	for key, value := range headers {
		w.Header()[key] = value
	}

	//header information
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(js))

	return nil
}

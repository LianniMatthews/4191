// Filename cmd/api/healthcheck.go
package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//js := `{"status":"available", "environment":%q, "version":%q}`
	data := envelope{
		"status": "available",
		"system_Info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		app.serveErrorResponse(w, r, err)
		return
	}
}

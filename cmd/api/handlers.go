// Filename cmd/api/schools.go
package main

import (
	"net/http"

	"github.com/LianniMatthews/4191/internal/data"
)

func (app *application) showCourseHandler(w http.ResponseWriter, r *http.Request) {
	code, err := app.readCodeParams(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	//fmt.Fprintf(w, "show details of %d\n", id)
	school := data.Course{
		Code:   code,
		Title:  "Advance Web Application",
		Credit: 3,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)

	if err != nil {
		app.logger.Println(err)
		app.serveErrorResponse(w, r, err)
		return
	}

}

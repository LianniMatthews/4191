//Filename cmd/api/routes.go

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	//New Router
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/courses", app.createCourseHandler)
	router.HandlerFunc(http.MethodGet, "/v1/courses/:id", app.showCourseHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/courses/:id", app.updateCourseHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/courses/:id", app.deleteCourseHandler)

	//return router
	return router
}

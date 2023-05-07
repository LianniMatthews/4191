// Filename cmd/api/schools.go
package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/LianniMatthews/4191/internal/data"
	"github.com/LianniMatthews/4191/internal/validator"
)

func (app *application) createCourseHandler(w http.ResponseWriter, r *http.Request) {
	//struct to hold course provided by request
	var input struct {
		Code   string `json:"code"`
		Title  string `json:"title"`
		Credit string  `json:"credit"`
	}
	//decode JSON request
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)

		return
	}
	// Copy the values from the input struct to the new Course struct
	course := &data.Course{
		Code:   input.Code,
		Title:  input.Title,
		Credit: input.Credit,
	}

	v := validator.New()
	// Check for validation errrors
	if data.ValidateCourse(v, course); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// write course to database
	err = app.models.Courses.Insert(course)
	if err != nil {
		app.serveErrorResponse(w, r, err)
		return
	}

	//creation header
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/schools/%d", course.ID))

	//write response
	err = app.writeJSON(w, http.StatusCreated, envelope{"course": course}, headers)
	if err != nil {
		app.serveErrorResponse(w, r, err)
		return
	}
}

func (app *application) showCourseHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Fetch the specific course
	course, err := app.models.Courses.Get(id)
	// Handle errors
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serveErrorResponse(w, r, err)
		}
		return
	}
	// Write the data returned by Get()
	err = app.writeJSON(w, http.StatusOK, envelope{"course": course}, nil)
	if err != nil {
		app.serveErrorResponse(w, r, err)
	}

}

func (app *application) updateCourseHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	//fetch original school
	course, err := app.models.Courses.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serveErrorResponse(w, r, err)
		}
		return
	}

	//course
	var input struct {
		Code   *string `json:"code"`
		Title  *string `json:"title"`
		Credit *string `json:"credit"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)

		return
	}

	//check for updates
	if input.Code != nil {
		course.Code = *input.Code
	}
	if input.Title != nil {
		course.Title = *input.Title
	}
	if input.Code != nil {
		course.Code = *input.Code
	}
	

	//validate
	v := validator.New()
	// Check for validation errrors
	if data.ValidateCourse(v, course); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	//update
	err = app.models.Courses.Update(course)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serveErrorResponse(w, r, err)
		}
		return
	}

	//write response
	err = app.writeJSON(w, http.StatusOK, envelope{"course": course}, nil)
	if err != nil {
		app.serveErrorResponse(w, r, err)
		return
	}
}

func (app *application) deleteCourseHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Courses.Delete(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serveErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "school succesfully deleted"}, nil)

	if err != nil {
		app.serveErrorResponse(w, r, err)
	}

}
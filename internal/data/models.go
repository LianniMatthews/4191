package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict = errors.New("edit conflict")
)

// Wrapper for Models
type Models struct {
	Courses CourseModel
}

// Create a Models instance
func NewModels(db *sql.DB) Models {
	return Models{
		Courses: CourseModel{DB: db},
	}
}

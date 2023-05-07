package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/LianniMatthews/4191/internal/validator"
)

// course represents one row of data in database
type Course struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Title     string    `json:"title"`
	Credit    string    `json:"credit"`
	CreatedAt time.Time `json:"-"`
	Version   int32     `json:"version"`
}

func ValidateCourse(v *validator.Validator, course *Course) {
	// Use the Check() method to execute our validation checks
	v.Check(course.Code != "", "code", "must be provided")
	v.Check(len(course.Code) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(course.Title != "", "title", "must be provided")
	v.Check(len(course.Title) <= 200, "level", "must not be more than 200 bytes long")

	v.Check(course.Credit != "", "credit", "must be provided")
	v.Check(len(course.Credit) <= 200, "level", "must not be more than 200 bytes long")
}

// Models
type CourseModel struct {
	DB *sql.DB
}

// Insert
func (m CourseModel) Insert(course *Course) error {
	//SQL Insert
	query := `
				INSERT INTO courses(code, title, credit)
				VALUES($1, $2, $3)
				RETURNING id, created_at, version
	`
	// Place data fields into a slice
	args := []interface{}{course.Code, course.Title, course.Credit}

	return m.DB.QueryRow(query, args...).Scan(&course.ID, &course.CreatedAt, &course.Version)
}

// Get
func (m CourseModel) Get(id int64) (*Course, error) {
	// validate id
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// SQL Select (get)
	query := `
		SELECT id, created_at, code, title, credit, version
		FROM courses
		WHERE id = $1
	`
	// course variable
	var course Course
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup
	defer cancel()
	// Execute the query
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&course.ID,
		&course.CreatedAt,
		&course.Code,
		&course.Title,
		&course.Credit,
		&course.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &course, nil
}

// Update ((AND version = $10))
func (m CourseModel) Update(course *Course) error {
	query := `
		UPDATE courses
		SET code = $1, title = $2, credit = $3, version = version + 1
		WHERE id = $4
		AND version = $5
		RETURNING version
	`
	args := []interface{}{
		course.Code,
		course.Title,
		course.Credit,
		course.ID,
		course.Version,
	}

	//check for edit conflicts
	err := m.DB.QueryRow(query, args...).Scan(&course.Version)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

// Delete
func (m CourseModel) Delete(id int64) error {
	//check id
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
			DELETE FROM courses
			WHERE id = $1
	`
	//Execute the query
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	//check affected rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

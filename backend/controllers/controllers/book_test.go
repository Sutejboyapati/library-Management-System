package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"library/backend/config"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllBooksReturnsList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to init sqlmock: %v", err)
	}
	defer db.Close()
	config.DB = db

	rows := sqlmock.NewRows([]string{"id", "title", "author", "isbn", "genre", "language", "shelf_number", "available_copies"}).
		AddRow(1, "Clean Code", "Robert C. Martin", "9780132350884", "Programming", "English", "A-1", 3)

	mock.ExpectQuery("SELECT id, title, author, isbn, COALESCE\\(genre,''\\), COALESCE\\(language,''\\), COALESCE\\(shelf_number,''\\), available_copies FROM Books").
		WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/api/books?title=Clean", nil)
	rr := httptest.NewRecorder()

	GetAllBooks(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	if rr.Body.String() == "" {
		t.Fatalf("expected json response body")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestGetAllBooksReturnsServerErrorOnDBFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to init sqlmock: %v", err)
	}
	defer db.Close()
	config.DB = db

	mock.ExpectQuery("SELECT id, title, author, isbn, COALESCE\\(genre,''\\), COALESCE\\(language,''\\), COALESCE\\(shelf_number,''\\), available_copies FROM Books").
		WillReturnError(assertAnError{})

	req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
	rr := httptest.NewRecorder()

	GetAllBooks(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", rr.Code)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

type assertAnError struct{}

func (assertAnError) Error() string { return "db error" }

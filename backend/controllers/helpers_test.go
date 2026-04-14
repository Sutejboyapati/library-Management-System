package controllers

import (
	"testing"

	"library/backend/models"
)

func TestNormalizePagination(t *testing.T) {
	page, limit, offset := normalizePagination(0, 150)
	if page != 1 || limit != 100 || offset != 0 {
		t.Fatalf("unexpected pagination values: page=%d limit=%d offset=%d", page, limit, offset)
	}
}

func TestValidateCredentials(t *testing.T) {
	if err := validateCredentials("reader", "secret1"); err != nil {
		t.Fatalf("expected credentials to be valid, got %v", err)
	}
	if err := validateCredentials("", "secret1"); err == nil {
		t.Fatal("expected empty username to fail validation")
	}
	if err := validateCredentials("reader", "123"); err == nil {
		t.Fatal("expected short password to fail validation")
	}
}

func TestValidateBookPayload(t *testing.T) {
	validBook := models.Book{Title: "Clean Code", Author: "Robert C. Martin", AvailableCopies: 2}
	if err := validateBookPayload(validBook); err != nil {
		t.Fatalf("expected book payload to be valid, got %v", err)
	}
	if err := validateBookPayload(models.Book{Author: "Author", AvailableCopies: 1}); err == nil {
		t.Fatal("expected missing title to fail validation")
	}
	if err := validateBookPayload(models.Book{Title: "Title", Author: "Author", AvailableCopies: -1}); err == nil {
		t.Fatal("expected negative copies to fail validation")
	}
}

func TestValidateBorrowRequest(t *testing.T) {
	if err := validateBorrowRequest(10); err != nil {
		t.Fatalf("expected borrow request to be valid, got %v", err)
	}
	if err := validateBorrowRequest(0); err == nil {
		t.Fatal("expected invalid book id to fail validation")
	}
}

func TestValidateReviewPayload(t *testing.T) {
	if err := validateReviewPayload(5, "Very useful and easy to read."); err != nil {
		t.Fatalf("expected review payload to be valid, got %v", err)
	}
	if err := validateReviewPayload(0, "Bad rating"); err == nil {
		t.Fatal("expected rating below 1 to fail validation")
	}
	if err := validateReviewPayload(4, "   "); err == nil {
		t.Fatal("expected blank comment to fail validation")
	}
}

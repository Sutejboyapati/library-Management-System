package controllers

import (
	"errors"
	"strings"

	"library/backend/models"
)

func normalizePagination(page, limit int) (int, int, int) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 50
	}
	if limit > 100 {
		limit = 100
	}
	return page, limit, (page - 1) * limit
}

func normalizeSearchTerm(search string) string {
	return strings.TrimSpace(search)
}

func validateCredentials(username, password string) error {
	if strings.TrimSpace(username) == "" {
		return errors.New("username is required")
	}
	if strings.TrimSpace(password) == "" {
		return errors.New("password is required")
	}
	if len(strings.TrimSpace(password)) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}

func validateBookPayload(book models.Book) error {
	if strings.TrimSpace(book.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(book.Author) == "" {
		return errors.New("author is required")
	}
	if book.AvailableCopies < 0 {
		return errors.New("available_copies cannot be negative")
	}
	return nil
}

func validateBorrowRequest(bookID int) error {
	if bookID < 1 {
		return errors.New("bookId must be a positive integer")
	}
	return nil
}

func validateReviewPayload(rating int, comment string) error {
	if rating < 1 || rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	if strings.TrimSpace(comment) == "" {
		return errors.New("comment is required")
	}
	if len(strings.TrimSpace(comment)) > 500 {
		return errors.New("comment must be 500 characters or fewer")
	}
	return nil
}

func sanitizeReviewComment(comment string) string {
	return strings.TrimSpace(comment)
}

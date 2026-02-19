package models

import (
	"database/sql"
	"time"
)

type BorrowRequest struct {
	UserID int `json:"userId"`
	BookID int `json:"bookId"`
}

type BorrowingRecord struct {
	ID         int           `json:"id"`
	UserID     int           `json:"user_id"`
	BookID     int           `json:"book_id"`
	Title      string        `json:"title"`
	Author     string        `json:"author"`
	ISBN       string        `json:"isbn"`
	BorrowedAt time.Time     `json:"borrowed_at"`
	DueDate    time.Time     `json:"due_date"`
	ReturnedAt sql.NullTime  `json:"returned_at,omitempty"`
	Status     string        `json:"status"`
}

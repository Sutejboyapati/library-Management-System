package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"library/backend/config"
	"library/backend/models"
)

func BorrowBook(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found", http.StatusUnauthorized)
		return
	}
	var req models.BorrowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	tx, err := config.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	var available int
	err = tx.QueryRow("SELECT available_copies FROM Books WHERE id = ?", req.BookID).Scan(&available)
	if err == sql.ErrNoRows {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("BorrowBook:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if available < 1 {
		http.Error(w, "No copies available", http.StatusBadRequest)
		return
	}

	rows, err := tx.Exec("UPDATE Books SET available_copies = available_copies - 1 WHERE id = ? AND available_copies > 0", req.BookID)
	if err != nil {
		log.Println("BorrowBook update:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	n, _ := rows.RowsAffected()
	if n == 0 {
		http.Error(w, "No copies available", http.StatusBadRequest)
		return
	}

	borrowedAt := time.Now()
	dueDate := borrowedAt.AddDate(0, 0, 14)
	_, err = tx.Exec("INSERT INTO borrowingrecords (user_id, book_id, borrowed_at, due_date) VALUES (?, ?, ?, ?)",
		userID, req.BookID, borrowedAt, dueDate)
	if err != nil {
		log.Println("BorrowBook insert:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if err := tx.Commit(); err != nil {
		http.Error(w, "Transaction error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book borrowed successfully"})
}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found", http.StatusUnauthorized)
		return
	}
	var req models.BorrowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	tx, err := config.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	var recordID int
	err = tx.QueryRow(
		"SELECT id FROM borrowingrecords WHERE user_id = ? AND book_id = ? AND returned_at IS NULL ORDER BY borrowed_at DESC LIMIT 1",
		userID, req.BookID,
	).Scan(&recordID)
	if err == sql.ErrNoRows {
		http.Error(w, "No active borrow record found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("ReturnBook:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	returnedAt := time.Now()
	_, err = tx.Exec("UPDATE borrowingrecords SET returned_at = ? WHERE id = ?", returnedAt, recordID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	_, err = tx.Exec("UPDATE Books SET available_copies = available_copies + 1 WHERE id = ?", req.BookID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if err := tx.Commit(); err != nil {
		http.Error(w, "Transaction error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book returned successfully"})
}

package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

<<<<<<< HEAD
	"library/backend/apiutil"
=======
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
	"library/backend/config"
	"library/backend/models"
)

func BorrowBook(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusUnauthorized, "User ID not found")
=======
		http.Error(w, "User ID not found", http.StatusUnauthorized)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	var req models.BorrowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid request data")
		return
	}
	if err := validateBorrowRequest(req.BookID); err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, err.Error())
=======
		http.Error(w, "Invalid request data", http.StatusBadRequest)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}

	tx, err := config.DB.Begin()
	if err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Failed to start transaction")
=======
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	defer tx.Rollback()

	var available int
	err = tx.QueryRow("SELECT available_copies FROM Books WHERE id = ?", req.BookID).Scan(&available)
	if err == sql.ErrNoRows {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusNotFound, "Book not found")
=======
		http.Error(w, "Book not found", http.StatusNotFound)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	if err != nil {
		log.Println("BorrowBook:", err)
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if available < 1 {
		apiutil.WriteError(w, http.StatusBadRequest, "No copies available")
		return
	}

	var activeBorrowCount int
	err = tx.QueryRow(
		"SELECT COUNT(*) FROM borrowingrecords WHERE user_id = ? AND book_id = ? AND returned_at IS NULL",
		userID, req.BookID,
	).Scan(&activeBorrowCount)
	if err != nil {
		log.Println("BorrowBook activeBorrowCount:", err)
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if activeBorrowCount > 0 {
		apiutil.WriteError(w, http.StatusBadRequest, "You already have an active borrowing for this book")
=======
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if available < 1 {
		http.Error(w, "No copies available", http.StatusBadRequest)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}

	rows, err := tx.Exec("UPDATE Books SET available_copies = available_copies - 1 WHERE id = ? AND available_copies > 0", req.BookID)
	if err != nil {
		log.Println("BorrowBook update:", err)
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
=======
		http.Error(w, "Database error", http.StatusInternalServerError)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	n, _ := rows.RowsAffected()
	if n == 0 {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusBadRequest, "No copies available")
=======
		http.Error(w, "No copies available", http.StatusBadRequest)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}

	borrowedAt := time.Now()
	dueDate := borrowedAt.AddDate(0, 0, 14)
	_, err = tx.Exec("INSERT INTO borrowingrecords (user_id, book_id, borrowed_at, due_date) VALUES (?, ?, ?, ?)",
		userID, req.BookID, borrowedAt, dueDate)
	if err != nil {
		log.Println("BorrowBook insert:", err)
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if err := tx.Commit(); err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, "Transaction error")
		return
	}

	apiutil.WriteJSON(w, http.StatusOK, map[string]string{"message": "Book borrowed successfully"})
=======
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if err := tx.Commit(); err != nil {
		http.Error(w, "Transaction error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book borrowed successfully"})
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusUnauthorized, "User ID not found")
=======
		http.Error(w, "User ID not found", http.StatusUnauthorized)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	var req models.BorrowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid request data")
		return
	}
	if err := validateBorrowRequest(req.BookID); err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, err.Error())
=======
		http.Error(w, "Invalid request data", http.StatusBadRequest)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}

	tx, err := config.DB.Begin()
	if err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Failed to start transaction")
=======
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	defer tx.Rollback()

	var recordID int
	err = tx.QueryRow(
		"SELECT id FROM borrowingrecords WHERE user_id = ? AND book_id = ? AND returned_at IS NULL ORDER BY borrowed_at DESC LIMIT 1",
		userID, req.BookID,
	).Scan(&recordID)
	if err == sql.ErrNoRows {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusNotFound, "No active borrow record found")
=======
		http.Error(w, "No active borrow record found", http.StatusNotFound)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	if err != nil {
		log.Println("ReturnBook:", err)
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
=======
		http.Error(w, "Database error", http.StatusInternalServerError)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}

	returnedAt := time.Now()
	_, err = tx.Exec("UPDATE borrowingrecords SET returned_at = ? WHERE id = ?", returnedAt, recordID)
	if err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
=======
		http.Error(w, "Database error", http.StatusInternalServerError)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	_, err = tx.Exec("UPDATE Books SET available_copies = available_copies + 1 WHERE id = ?", req.BookID)
	if err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if err := tx.Commit(); err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, "Transaction error")
		return
	}

	apiutil.WriteJSON(w, http.StatusOK, map[string]string{"message": "Book returned successfully"})
=======
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if err := tx.Commit(); err != nil {
		http.Error(w, "Transaction error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book returned successfully"})
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
}

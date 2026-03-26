package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"library/backend/config"
	"library/backend/models"

	"github.com/gorilla/mux"
)

func GetUserBorrowings(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	tokenUserID, _ := r.Context().Value("userID").(int)
	role, _ := r.Context().Value("role").(string)
	if tokenUserID != userID && role != "admin" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	query := `
		SELECT b.id, b.title, b.author, b.isbn, br.borrowed_at, br.due_date, br.returned_at
		FROM borrowingrecords br
		JOIN Books b ON br.book_id = b.id
		WHERE br.user_id = ?
		ORDER BY br.borrowed_at DESC`
	rows, err := config.DB.Query(query, userID)
	if err != nil {
		log.Println("GetUserBorrowings:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var records []models.BorrowingRecord
	for rows.Next() {
		var rec models.BorrowingRecord
		err := rows.Scan(&rec.BookID, &rec.Title, &rec.Author, &rec.ISBN, &rec.BorrowedAt, &rec.DueDate, &rec.ReturnedAt)
		if err != nil {
			continue
		}
		if rec.ReturnedAt.Valid {
			rec.Status = "Returned"
		} else {
			rec.Status = "Borrowing"
		}
		records = append(records, rec)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

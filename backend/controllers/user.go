package controllers

import (
<<<<<<< HEAD
=======
	"encoding/json"
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
	"log"
	"net/http"
	"strconv"

<<<<<<< HEAD
	"library/backend/apiutil"
=======
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
	"library/backend/config"
	"library/backend/models"

	"github.com/gorilla/mux"
)

func GetUserBorrowings(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid user ID")
=======
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	tokenUserID, _ := r.Context().Value("userID").(int)
	role, _ := r.Context().Value("role").(string)
	if tokenUserID != userID && role != "admin" {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusForbidden, "Forbidden")
=======
		http.Error(w, "Forbidden", http.StatusForbidden)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
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
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
=======
		http.Error(w, "Database error", http.StatusInternalServerError)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
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

<<<<<<< HEAD
	apiutil.WriteJSON(w, http.StatusOK, records)
=======
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
}

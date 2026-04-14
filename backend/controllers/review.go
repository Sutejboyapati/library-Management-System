package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"library/backend/apiutil"
	"library/backend/config"
	"library/backend/models"

	"github.com/gorilla/mux"
)

func GetBookReviews(w http.ResponseWriter, r *http.Request) {
	bookID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	rows, err := config.DB.Query(`
		SELECT r.id, r.book_id, r.user_id, u.username, r.rating, r.comment, r.created_at, r.updated_at
		FROM Reviews r
		JOIN Users u ON u.id = r.user_id
		WHERE r.book_id = ?
		ORDER BY r.updated_at DESC, r.created_at DESC
	`, bookID)
	if err != nil {
		log.Println("GetBookReviews:", err)
		apiutil.WriteError(w, http.StatusInternalServerError, "Failed to load reviews")
		return
	}
	defer rows.Close()

	reviews := make([]models.Review, 0)
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(
			&review.ID,
			&review.BookID,
			&review.UserID,
			&review.Username,
			&review.Rating,
			&review.Comment,
			&review.CreatedAt,
			&review.UpdatedAt,
		); err != nil {
			log.Println("GetBookReviews scan:", err)
			continue
		}
		reviews = append(reviews, review)
	}

	apiutil.WriteJSON(w, http.StatusOK, reviews)
}

func UpsertBookReview(w http.ResponseWriter, r *http.Request) {
	bookID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		apiutil.WriteError(w, http.StatusUnauthorized, "User ID not found")
		return
	}

	var req models.ReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validateReviewPayload(req.Rating, req.Comment); err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	var exists bool
	err = config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM Books WHERE id = ?)", bookID).Scan(&exists)
	if err != nil {
		log.Println("UpsertBookReview exists:", err)
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if !exists {
		apiutil.WriteError(w, http.StatusNotFound, "Book not found")
		return
	}

	result, err := config.DB.Exec(`
		INSERT INTO Reviews (book_id, user_id, rating, comment)
		VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE rating = VALUES(rating), comment = VALUES(comment), updated_at = CURRENT_TIMESTAMP
	`, bookID, userID, req.Rating, sanitizeReviewComment(req.Comment))
	if err != nil {
		log.Println("UpsertBookReview insert:", err)
		apiutil.WriteError(w, http.StatusInternalServerError, "Failed to save review")
		return
	}

	status := http.StatusCreated
	if rows, _ := result.RowsAffected(); rows > 1 {
		status = http.StatusOK
	}

	var review models.Review
	err = config.DB.QueryRow(`
		SELECT r.id, r.book_id, r.user_id, u.username, r.rating, r.comment, r.created_at, r.updated_at
		FROM Reviews r
		JOIN Users u ON u.id = r.user_id
		WHERE r.book_id = ? AND r.user_id = ?
	`, bookID, userID).Scan(
		&review.ID,
		&review.BookID,
		&review.UserID,
		&review.Username,
		&review.Rating,
		&review.Comment,
		&review.CreatedAt,
		&review.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			apiutil.WriteJSON(w, status, map[string]string{"message": "Review saved successfully"})
			return
		}
		log.Println("UpsertBookReview fetch:", err)
		apiutil.WriteError(w, http.StatusInternalServerError, "Failed to load saved review")
		return
	}

	apiutil.WriteJSON(w, status, review)
}

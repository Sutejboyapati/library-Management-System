package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"library/backend/config"
	"library/backend/models"

	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("title")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 50
	}
	offset := (page - 1) * limit

	query := "SELECT id, title, author, isbn, COALESCE(genre,''), COALESCE(language,''), COALESCE(shelf_number,''), available_copies FROM Books"
	args := []interface{}{}
	if search != "" {
		query += " WHERE title LIKE ? OR author LIKE ? OR genre LIKE ?"
		like := "%" + search + "%"
		args = append(args, like, like, like)
	}
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		log.Println("GetAllBooks:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Genre, &b.Language, &b.ShelfNumber, &b.AvailableCopies)
		if err != nil {
			continue
		}
		books = append(books, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var b models.Book
	err = config.DB.QueryRow(
		"SELECT id, title, author, isbn, COALESCE(genre,''), COALESCE(language,''), COALESCE(shelf_number,''), available_copies FROM Books WHERE id = ?", id,
	).Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Genre, &b.Language, &b.ShelfNumber, &b.AvailableCopies)
	if err == sql.ErrNoRows {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("GetBookByID:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var b models.Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res, err := config.DB.Exec(
		"INSERT INTO Books (title, author, genre, language, shelf_number, available_copies, isbn) VALUES (?, ?, ?, ?, ?, ?, ?)",
		b.Title, b.Author, b.Genre, b.Language, b.ShelfNumber, b.AvailableCopies, b.ISBN,
	)
	if err != nil {
		log.Println("AddBook:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Book added successfully", "bookId": id})
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var b models.Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res, err := config.DB.Exec(
		"UPDATE Books SET title=?, author=?, genre=?, language=?, shelf_number=?, available_copies=?, isbn=? WHERE id=?",
		b.Title, b.Author, b.Genre, b.Language, b.ShelfNumber, b.AvailableCopies, b.ISBN, id,
	)
	if err != nil {
		log.Println("UpdateBook:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book updated successfully"})
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	res, err := config.DB.Exec("DELETE FROM Books WHERE id = ?", id)
	if err != nil {
		log.Println("DeleteBook:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted successfully"})
}

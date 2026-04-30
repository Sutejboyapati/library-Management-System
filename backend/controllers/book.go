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

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("q")
	if search == "" {
		search = r.URL.Query().Get("title")
	}
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	page, limit, offset := normalizePagination(page, limit)
	search = normalizeSearchTerm(search)

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
		apiutil.WriteError(w, http.StatusInternalServerError, "Server error")
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

	apiutil.WriteJSON(w, http.StatusOK, books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var b models.Book
	err = config.DB.QueryRow(
		"SELECT id, title, author, isbn, COALESCE(genre,''), COALESCE(language,''), COALESCE(shelf_number,''), available_copies FROM Books WHERE id = ?", id,
	).Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Genre, &b.Language, &b.ShelfNumber, &b.AvailableCopies)
	if err == sql.ErrNoRows {
		apiutil.WriteError(w, http.StatusNotFound, "Book not found")
		return
	}
	if err != nil {
		log.Println("GetBookByID:", err)
		apiutil.WriteError(w, http.StatusInternalServerError, "Server error")
		return
	}

	apiutil.WriteJSON(w, http.StatusOK, b)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var b models.Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validateBookPayload(b); err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := config.DB.Exec(
		"INSERT INTO Books (title, author, genre, language, shelf_number, available_copies, isbn) VALUES (?, ?, ?, ?, ?, ?, ?)",
		b.Title, b.Author, b.Genre, b.Language, b.ShelfNumber, b.AvailableCopies, b.ISBN,
	)
	if err != nil {
		log.Println("AddBook:", err)
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
		return
	}

	id, _ := res.LastInsertId()
	apiutil.WriteJSON(w, http.StatusCreated, map[string]interface{}{"message": "Book added successfully", "bookId": id})
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var b models.Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validateBookPayload(b); err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := config.DB.Exec(
		"UPDATE Books SET title=?, author=?, genre=?, language=?, shelf_number=?, available_copies=?, isbn=? WHERE id=?",
		b.Title, b.Author, b.Genre, b.Language, b.ShelfNumber, b.AvailableCopies, b.ISBN, id,
	)
	if err != nil {
		log.Println("UpdateBook:", err)
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		apiutil.WriteError(w, http.StatusNotFound, "Book not found")
		return
	}

	apiutil.WriteJSON(w, http.StatusOK, map[string]string{"message": "Book updated successfully"})
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	res, err := config.DB.Exec("DELETE FROM Books WHERE id = ?", id)
	if err != nil {
		log.Println("DeleteBook:", err)
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		apiutil.WriteError(w, http.StatusNotFound, "Book not found")
		return
	}

	apiutil.WriteJSON(w, http.StatusOK, map[string]string{"message": "Book deleted successfully"})
}

package controllers

import (
	"database/sql"
	"encoding/json"
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

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	search := r.URL.Query().Get("q")
	if search == "" {
		search = r.URL.Query().Get("title")
	}
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	page, limit, offset := normalizePagination(page, limit)
	search = normalizeSearchTerm(search)
=======
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
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504

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
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Server error")
=======
		http.Error(w, "Server error", http.StatusInternalServerError)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
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

<<<<<<< HEAD
	apiutil.WriteJSON(w, http.StatusOK, books)
=======
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid book ID")
=======
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}

	var b models.Book
	err = config.DB.QueryRow(
		"SELECT id, title, author, isbn, COALESCE(genre,''), COALESCE(language,''), COALESCE(shelf_number,''), available_copies FROM Books WHERE id = ?", id,
	).Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Genre, &b.Language, &b.ShelfNumber, &b.AvailableCopies)
	if err == sql.ErrNoRows {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusNotFound, "Book not found")
=======
		http.Error(w, "Book not found", http.StatusNotFound)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	if err != nil {
		log.Println("GetBookByID:", err)
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Server error")
		return
	}

	apiutil.WriteJSON(w, http.StatusOK, b)
=======
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var b models.Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validateBookPayload(b); err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, err.Error())
=======
		http.Error(w, "Invalid request body", http.StatusBadRequest)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}

	res, err := config.DB.Exec(
		"INSERT INTO Books (title, author, genre, language, shelf_number, available_copies, isbn) VALUES (?, ?, ?, ?, ?, ?, ?)",
		b.Title, b.Author, b.Genre, b.Language, b.ShelfNumber, b.AvailableCopies, b.ISBN,
	)
	if err != nil {
		log.Println("AddBook:", err)
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
=======
		http.Error(w, "Database error", http.StatusInternalServerError)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}

	id, _ := res.LastInsertId()
<<<<<<< HEAD
	apiutil.WriteJSON(w, http.StatusCreated, map[string]interface{}{"message": "Book added successfully", "bookId": id})
=======
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Book added successfully", "bookId": id})
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid book ID")
=======
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}

	var b models.Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validateBookPayload(b); err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, err.Error())
=======
		http.Error(w, "Invalid request body", http.StatusBadRequest)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}

	res, err := config.DB.Exec(
		"UPDATE Books SET title=?, author=?, genre=?, language=?, shelf_number=?, available_copies=?, isbn=? WHERE id=?",
		b.Title, b.Author, b.Genre, b.Language, b.ShelfNumber, b.AvailableCopies, b.ISBN, id,
	)
	if err != nil {
		log.Println("UpdateBook:", err)
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
=======
		http.Error(w, "Database error", http.StatusInternalServerError)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusNotFound, "Book not found")
		return
	}

	apiutil.WriteJSON(w, http.StatusOK, map[string]string{"message": "Book updated successfully"})
=======
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book updated successfully"})
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid book ID")
=======
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}

	res, err := config.DB.Exec("DELETE FROM Books WHERE id = ?", id)
	if err != nil {
		log.Println("DeleteBook:", err)
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
=======
		http.Error(w, "Database error", http.StatusInternalServerError)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusNotFound, "Book not found")
		return
	}

	apiutil.WriteJSON(w, http.StatusOK, map[string]string{"message": "Book deleted successfully"})
=======
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted successfully"})
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
}

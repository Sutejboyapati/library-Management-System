package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"library/backend/config"

	"golang.org/x/crypto/bcrypt"
)

// Open Library API response (subset we need)
type openLibrarySearch struct {
	Docs []struct {
		Title   string   `json:"title"`
		Authors []string `json:"author_name"`
		Subject []string `json:"subject"`
		ISBN    []string `json:"isbn"`
	} `json:"docs"`
}

// Seed inserts admin user and fetches real books from Open Library API. Safe to call multiple times.
// Add ?force=1 to clear existing books and reload from the internet.
func Seed(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	force := r.URL.Query().Get("force") == "1"
	if force {
		_, _ = config.DB.Exec("DELETE FROM borrowingrecords")
		_, _ = config.DB.Exec("DELETE FROM Books")
	}

	// Check if we already have books (skip if force)
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM Books").Scan(&count)
	if err != nil {
		log.Println("Seed check:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Data already seeded", "booksCount": count})
		return
	}

	// Insert admin user (password: admin123) - skip if exists
	hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	_, _ = config.DB.Exec("INSERT IGNORE INTO Users (username, password, role) VALUES (?, ?, ?)", "admin", string(hash), "admin")

	// Fetch books from Open Library API
	added := fetchAndInsertBooksFromOpenLibrary()
	if added == 0 {
		// Fallback: insert a few hardcoded books if API fails
		fallbackBooks := []struct {
			title  string
			author string
			genre  string
			copies int
		}{
			{"The Great Gatsby", "F. Scott Fitzgerald", "Fiction", 5},
			{"To Kill a Mockingbird", "Harper Lee", "Fiction", 4},
			{"1984", "George Orwell", "Dystopian", 6},
			{"Pride and Prejudice", "Jane Austen", "Romance", 3},
			{"Clean Code", "Robert C. Martin", "Programming", 5},
		}
		for _, b := range fallbackBooks {
			_, _ = config.DB.Exec(
				"INSERT INTO Books (title, author, genre, available_copies) VALUES (?, ?, ?, ?)",
				b.title, b.author, b.genre, b.copies,
			)
			added++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Books loaded from Open Library", "booksAdded": added})
}

func fetchAndInsertBooksFromOpenLibrary() int {
	queries := []string{"fiction", "programming", "science"}
	client := &http.Client{Timeout: 15 * time.Second}
	seen := make(map[string]bool)
	added := 0

	for _, q := range queries {
		url := "https://openlibrary.org/search.json?q=" + q + "&limit=12"
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			continue
		}
		req.Header.Set("User-Agent", "LibraryManagementSystem/1.0")

		resp, err := client.Do(req)
		if err != nil {
			log.Println("Open Library request:", err)
			continue
		}

		var result openLibrarySearch
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			continue
		}
		resp.Body.Close()

		for _, doc := range result.Docs {
			if doc.Title == "" || len(doc.Authors) == 0 {
				continue
			}
			key := doc.Title + "|" + doc.Authors[0]
			if seen[key] {
				continue
			}
			seen[key] = true
			author := doc.Authors[0]
			if len(author) > 255 {
				author = author[:252] + "..."
			}
			title := doc.Title
			if len(title) > 500 {
				title = title[:497] + "..."
			}
			genre := "General"
			if len(doc.Subject) > 0 {
				genre = doc.Subject[0]
				if len(genre) > 100 {
					genre = genre[:100]
				}
			}
			isbn := ""
			if len(doc.ISBN) > 0 {
				isbn = doc.ISBN[0]
			}

			_, err := config.DB.Exec(
				"INSERT INTO Books (title, author, isbn, genre, available_copies) VALUES (?, ?, ?, ?, ?)",
				title, author, isbn, genre, 3,
			)
			if err != nil {
				if strings.Contains(err.Error(), "Duplicate") {
					continue
				}
				log.Println("Insert book:", err)
				continue
			}
			added++
		}
	}

	return added
}

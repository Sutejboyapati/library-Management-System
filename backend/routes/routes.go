package routes

import (
	"net/http"

	"library/backend/controllers"
	"library/backend/middleware"

	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/register", controllers.Register).Methods(http.MethodPost)
	api.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)
	api.HandleFunc("/seed", controllers.Seed).Methods(http.MethodGet, http.MethodPost)
	api.HandleFunc("/dashboard/summary", controllers.GetDashboardSummary).Methods(http.MethodGet)

	api.HandleFunc("/books", controllers.GetAllBooks).Methods(http.MethodGet)
	api.HandleFunc("/books/{id:[0-9]+}", controllers.GetBookByID).Methods(http.MethodGet)
	api.HandleFunc("/books/{id:[0-9]+}/reviews", controllers.GetBookReviews).Methods(http.MethodGet)
	api.Handle("/books/{id:[0-9]+}/reviews", middleware.VerifyToken(http.HandlerFunc(controllers.UpsertBookReview))).Methods(http.MethodPost)
	api.Handle("/books", middleware.VerifyToken(middleware.RequireAdmin(http.HandlerFunc(controllers.AddBook)))).Methods(http.MethodPost)
	api.Handle("/books/{id:[0-9]+}", middleware.VerifyToken(middleware.RequireAdmin(http.HandlerFunc(controllers.UpdateBook)))).Methods(http.MethodPut)
	api.Handle("/books/{id:[0-9]+}", middleware.VerifyToken(middleware.RequireAdmin(http.HandlerFunc(controllers.DeleteBook)))).Methods(http.MethodDelete)

	api.Handle("/borrow", middleware.VerifyToken(http.HandlerFunc(controllers.BorrowBook))).Methods(http.MethodPost)
	api.Handle("/borrow/return", middleware.VerifyToken(http.HandlerFunc(controllers.ReturnBook))).Methods(http.MethodPost)

	api.Handle("/users/{id:[0-9]+}/borrowings", middleware.VerifyToken(http.HandlerFunc(controllers.GetUserBorrowings))).Methods(http.MethodGet)
}

package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"library/backend/apiutil"
	"library/backend/config"
)

type DashboardSummary struct {
	TotalBooks       int `json:"totalBooks"`
	AvailableBooks   int `json:"availableBooks"`
	ActiveBorrowings int `json:"activeBorrowings"`
}

func GetDashboardSummary(w http.ResponseWriter, r *http.Request) {
	summary := DashboardSummary{}

	if err := config.DB.QueryRow("SELECT COUNT(*) FROM Books").Scan(&summary.TotalBooks); err != nil {
		log.Println("GetDashboardSummary totalBooks:", err)
		apiutil.WriteError(w, http.StatusInternalServerError, "Failed to load dashboard summary")
		return
	}

	if err := config.DB.QueryRow("SELECT COALESCE(SUM(available_copies), 0) FROM Books").Scan(&summary.AvailableBooks); err != nil {
		log.Println("GetDashboardSummary availableBooks:", err)
		apiutil.WriteError(w, http.StatusInternalServerError, "Failed to load dashboard summary")
		return
	}

	err := config.DB.QueryRow("SELECT COUNT(*) FROM borrowingrecords WHERE returned_at IS NULL").Scan(&summary.ActiveBorrowings)
	if err != nil && err != sql.ErrNoRows {
		log.Println("GetDashboardSummary activeBorrowings:", err)
		apiutil.WriteError(w, http.StatusInternalServerError, "Failed to load dashboard summary")
		return
	}

	apiutil.WriteJSON(w, http.StatusOK, summary)
}

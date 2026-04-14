package models

import "time"

type Review struct {
	ID        int       `json:"id"`
	BookID    int       `json:"bookId"`
	UserID    int       `json:"userId"`
	Username  string    `json:"username"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ReviewRequest struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

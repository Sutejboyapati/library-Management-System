package models

type Book struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	ISBN            string `json:"isbn"`
	Genre           string `json:"genre"`
	Language        string `json:"language"`
	ShelfNumber     string `json:"shelf_number"`
	AvailableCopies int    `json:"available_copies"`
}

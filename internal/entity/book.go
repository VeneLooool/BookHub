package entity

type Book struct {
	ID          int64  `json:"book_id" db:"book_id"`
	Title       string `json:"title" db:"title"`
	Author      string `json:"author" db:"author"`
	NumberPages int64  `json:"number_pages" db:"number_pages"`
	CurrentPage int64  `json:"current_page" db:"current_page"`
	Desc        string `json:"desc" db:"desc"`
	Image       File   `json:"image" db:"image"`
	File        File   `json:"file" db:"file"`
}

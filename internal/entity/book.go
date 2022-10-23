package entity

type Book struct {
	ID          int64
	Title       string
	Author      string
	NumberPages int64
	CurrentPage int64
	Desc        string
	Image       File
	File        File
}

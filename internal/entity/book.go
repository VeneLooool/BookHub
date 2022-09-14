package entity

import "image"

type Book struct {
	ID          int64
	Title       string
	Author      string
	NumberPages int64
	Desc        string
	Image       image.NRGBA
	File        []byte
}

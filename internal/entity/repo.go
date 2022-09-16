package entity

type Repo struct {
	ID         int64
	Visibility string
	Name       string
	Desc       string
	Books      []int64
}

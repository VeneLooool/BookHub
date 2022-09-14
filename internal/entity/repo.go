package entity

type Repo struct {
	ID      int64
	Visible bool
	Name    string
	Desc    string
	Books   []int64
	UsersID []int64
}

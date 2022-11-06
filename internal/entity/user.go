package entity

type User struct {
	ID       int64  `json:"user_id" db:"user_id"`
	Name     string `json:"name" db:"name"`
	UserName string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Desc     string `json:"user_desc" db:"user_desc"`
}

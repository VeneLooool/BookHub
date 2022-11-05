package entity

type Repo struct {
	ID         int64  `json:"repo_id" db:"repo_id"`
	Name       string `json:"name" db:"name"`
	Visibility string `json:"visible" db:"visible"`
	Desc       string `json:"repo_desc" db:"repo_desc"`
	UserID     int64  `json:"user_id" db:"user_id"`
}

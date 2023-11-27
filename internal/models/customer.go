package models

type Customer struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

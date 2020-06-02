package model

type User struct {
	Id       string `json:"id"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

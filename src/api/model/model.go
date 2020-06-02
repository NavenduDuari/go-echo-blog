package model

type UserJWT struct {
	Id    string `json:"id"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type UserForLogin struct {
	UserId   string `json:"userid"`
	Password string `json:"password"`
}

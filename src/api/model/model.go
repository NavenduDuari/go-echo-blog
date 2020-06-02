package model

import blogModel "github.com/NavenduDuari/go-echo-blog/src/db/blog/model"

type UserJWT struct {
	Id    string `json:"id"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type UserForLogin struct {
	UserId   string `json:"userid"`
	Password string `json:"password"`
}

type InputForBlogUpdate struct {
	Id   string         `json:"id"`
	Blog blogModel.Blog `json:"blog"`
}

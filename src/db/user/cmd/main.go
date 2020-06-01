package main

import (
	"github.com/NavenduDuari/go-echo-blog/src/db/user"
)

func main() {
	user.InitStore()
	// user.InsertUser(&model.User{})
	// row, _ := user.FetchUserById("dc5ed9ca-41f7-11ea-8d0b-122d51db1d75")
	// fmt.Println(row)
	// user.DeleteUser("dc5ed9ca-41f7-11ea-8d0b-122d51db1d75")
	user.CloseStore()
}

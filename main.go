package main

import (
	"github.com/NavenduDuari/go-echo-blog/src/db/auth"
	"github.com/NavenduDuari/go-echo-blog/src/db/blog"
	"github.com/NavenduDuari/go-echo-blog/src/db/user"
	"github.com/NavenduDuari/go-echo-blog/src/router"
)

func main() {

	e := router.New()

	// init database
	user.InitStore()
	defer user.CloseStore()
	blog.InitStore()
	defer blog.CloseStore()
	auth.InitStore()
	defer auth.CloseStore()

	e.Start(":8000")
}

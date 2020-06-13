package main

import (
	"github.com/NavenduDuari/go-echo-blog/src/db/blog"
	"github.com/NavenduDuari/go-echo-blog/src/db/otp"
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
	otp.InitStore()
	defer otp.CloseStore()

	e.Start(":8000")
}

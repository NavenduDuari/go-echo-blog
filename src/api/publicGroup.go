package api

import (
	"github.com/NavenduDuari/go-echo-blog/src/api/handlers"
	"github.com/labstack/echo"
)

func PublicGroup(e *echo.Group) {
	e.GET("/login", handlers.Login)
	e.POST("/signup", handlers.AddUser)
}

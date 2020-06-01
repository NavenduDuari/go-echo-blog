package router

import (
	"github.com/NavenduDuari/go-echo-blog/src/api"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// router groups
	publicGroup := e.Group("/public")
	protectedGroup := e.Group("/private")

	// set group routes
	api.PublicGroup(publicGroup)
	api.ProtectedGroup(protectedGroup)
	return e
}

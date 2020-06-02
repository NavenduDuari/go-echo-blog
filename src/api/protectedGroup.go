package api

import (
	"github.com/NavenduDuari/go-echo-blog/src/api/handlers"
	"github.com/labstack/echo"
)

func ProtectedGroup(e *echo.Group) {
	e.POST("/blogs", handlers.PostBlog)
	e.PUT("/blogs", handlers.EditBlog)
	e.DELETE("/blogs/:id", handlers.DeleteBlog)
}

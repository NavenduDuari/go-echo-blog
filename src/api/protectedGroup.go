package api

import (
	"github.com/NavenduDuari/go-echo-blog/src/api/handlers"
	"github.com/labstack/echo"
)

func ProtectedGroup(e *echo.Group) {
	e.GET("/blogs", handlers.GetBlogs)
	e.GET("/blog", handlers.GetBlogById)
	e.PUT("/blogs", handlers.EditBlog)
	e.DELETE("/blogs", handlers.DeleteBlog)
}

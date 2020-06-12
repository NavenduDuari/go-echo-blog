package api

import (
	"github.com/NavenduDuari/go-echo-blog/src/api/handlers"
	"github.com/labstack/echo"
)

func PublicGroup(e *echo.Group) {
	e.GET("/login", handlers.Login)
	e.GET("/login-with-otp", handlers.SendOtp)
	e.POST("/login-with-otp", handlers.VerifyOtp)
	e.POST("/signup", handlers.SignUp)
	e.GET("/blogs", handlers.GetBlogs)
	e.GET("/blog", handlers.GetBlogById)
}

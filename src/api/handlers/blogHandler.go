package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/NavenduDuari/go-echo-blog/src/db/blog"
	blogModel "github.com/NavenduDuari/go-echo-blog/src/db/blog/model"
	"github.com/labstack/echo"
)

func GetBlogs(c echo.Context) error {
	fmt.Println("get blog")
	return errors.New("")
}

func GetBlogById(c echo.Context) error {
	return errors.New("")
}

func PostBlog(c echo.Context) error {
	newBlog := new(blogModel.Blog)
	if err := c.Bind(newBlog); err != nil {
		return err
	}
	fmt.Println(newBlog)
	if err := blog.InsertBlog(newBlog); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Blog POST Failed")
	}

	return c.JSON(http.StatusCreated, "Blog POST successful")
}

func EditBlog(c echo.Context) error {
	return errors.New("")
}

func DeleteBlog(c echo.Context) error {
	return errors.New("")
}

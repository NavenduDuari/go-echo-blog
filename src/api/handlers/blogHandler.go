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
	blogs, err := blog.FetchBlogs()
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusForbidden, "Blog FETCH Failed")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Blogs FETCH successful!",
		"blogs":   blogs,
	})
}

func GetBlogById(c echo.Context) error {
	return errors.New("")
}

func PostBlog(c echo.Context) error {
	newBlog := new(blogModel.Blog)
	if err := c.Bind(newBlog); err != nil {
		return err
	}

	if err := blog.InsertBlog(newBlog); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusForbidden, "Blog POST Failed")
	}

	return c.JSON(http.StatusOK, "Blog POST successful")
}

func EditBlog(c echo.Context) error {
	return errors.New("")
}

func DeleteBlog(c echo.Context) error {
	id := c.Param("id")

	if err := blog.DeleteBlog(id); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusForbidden, "Blog DELETE Failed")
	}

	return c.JSON(http.StatusOK, "Blog DELETE successful")
}

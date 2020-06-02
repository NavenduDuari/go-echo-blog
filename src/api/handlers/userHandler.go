package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/NavenduDuari/go-echo-blog/src/db/user"
	userModel "github.com/NavenduDuari/go-echo-blog/src/db/user/model"
	"github.com/labstack/echo"
)

func AddUser(c echo.Context) error {
	newUser := new(userModel.User)
	if err := c.Bind(newUser); err != nil {
		return err
	}
	fmt.Println(newUser)
	if err := user.InsertUser(newUser); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusForbidden, "Failed SignUp")
	}

	return c.JSON(http.StatusOK, "SignUp successful")
}

func EditUser(c echo.Context) error {
	return errors.New("")
}

func DeleteUser(c echo.Context) error {
	return errors.New("")
}

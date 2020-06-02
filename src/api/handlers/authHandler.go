package handlers

import (
	"fmt"
	"net/http"

	"github.com/NavenduDuari/go-echo-blog/src/api/model"
	"github.com/NavenduDuari/go-echo-blog/src/common/utils"
	"github.com/NavenduDuari/go-echo-blog/src/db/user"
	userModel "github.com/NavenduDuari/go-echo-blog/src/db/user/model"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	userForLogin := new(userModel.User)
	if err := c.Bind(userForLogin); err != nil {
		return err
	}
	fmt.Println(userForLogin)
	userData, err := user.FetchUserById(userForLogin.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, "Failed Login")
	}

	if userData.Password == userForLogin.Password {
		token, err := utils.GetJwtLoginToken(model.UserJWT{
			Id:    userData.Id,
			Phone: userData.Phone,
			Email: userData.Email,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusNotAcceptable, "Failed Login")
		}
		return c.JSON(http.StatusCreated, map[string]string{
			"message": "Login successful!",
			"token":   token,
		})
	}

	return c.String(http.StatusUnauthorized, "Login failed! Password didn't match")
}

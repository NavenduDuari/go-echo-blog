package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	fmt.Println(c)
	// username := c.QueryParam("username")
	// password := c.QueryParam("password")

	// // check username and password against DB after hashing the password
	// if username == "jack" && password == "1234" {
	// 	cookie := &http.Cookie{}

	// 	// this is the same
	// 	//cookie := new(http.Cookie)

	// 	cookie.Name = "sessionID"
	// 	cookie.Value = "some_string"
	// 	cookie.Expires = time.Now().Add(48 * time.Hour)

	// 	c.SetCookie(cookie)

	// 	// create jwt token
	// 	token, err := utils.GetJwtLoginToken(model.UserJWT{})
	// 	if err != nil {
	// 		log.Println("Error Creating JWT token", err)
	// 		return c.String(http.StatusInternalServerError, "something went wrong")
	// 	}

	// 	return c.JSON(http.StatusOK, map[string]string{
	// 		"message": "You were logged in!",
	// 		"token":   token,
	// 	})
	// }

	return c.String(http.StatusUnauthorized, "Your username or password were wrong")
}

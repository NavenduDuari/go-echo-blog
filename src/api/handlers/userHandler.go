package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

func AddUser(c echo.Context) error {
	// user := models.User{}

	// defer c.Request().Body.Close()

	// rules := govalidator.MapData{
	// 	"name":  []string{"required"},
	// 	"email": []string{"required", "email"},
	// }

	// vld := ValidateRequest(c, rules, &user)
	// if vld != nil {
	// 	return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	// }

	// result, err := models.Create(&user)
	// if err != nil {
	// 	log.Printf("FAILED TO CREATE : %s\n", err)
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Failed to create new user")
	// }

	// return c.JSON(http.StatusCreated, result)
	return c.JSON(http.StatusUnauthorized, "this is from signup")
}

func EditUser(c echo.Context) error {
	// id, err := strconv.Atoi(c.QueryParam("id"))

	// defer c.Request().Body.Close()

	// rules := govalidator.MapData{
	// 	"name":  []string{},
	// 	"email": []string{"email"},
	// }

	// user, err := models.FindUserByID(id)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusNotFound, err.Error())
	// }

	// vld := ValidateRequest(c, rules, &user)
	// if vld != nil {
	// 	return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	// }

	// c.Bind(&user)

	// err = user.Update()
	// if err != nil {
	// 	log.Printf("FAILED TO UPDATE: %s\n", err)
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Failed to update user")
	// }

	// return c.JSON(http.StatusOK, user)
	return errors.New("")
}

func DeleteUser(c echo.Context) error {
	// id, err := strconv.Atoi(c.QueryParam("id"))

	// defer c.Request().Body.Close()

	// rules := govalidator.MapData{
	// 	"id": []string{"required", "numeric"},
	// }

	// vld := ValidateQueryStr(c, rules)
	// if vld != nil {
	// 	return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	// }

	// user, err := models.FindUserByID(id)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusNotFound, err.Error())
	// }

	// c.Bind(&user)

	// err = user.Delete()
	// if err != nil {
	// 	log.Printf("FAILED TO DELETE: %s\n", err)
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Failed to delete user")
	// }

	// return c.JSON(http.StatusOK, user)
	return errors.New("")
}

package handlers

import (
	"errors"
	"fmt"

	"github.com/labstack/echo"
)

func GetBlogs(c echo.Context) error {
	fmt.Println(c)
	// rp, err := strconv.Atoi(c.QueryParam("rp"))
	// page, err := strconv.Atoi(c.QueryParam("p"))
	// name := c.QueryParam("name")
	// email := c.QueryParam("email")

	// defer c.Request().Body.Close()

	// rules := govalidator.MapData{
	// 	"rp":    []string{"numeric"},
	// 	"page":  []string{"numeric"},
	// 	"name":  []string{"alpha_num"},
	// 	"email": []string{"email"},
	// }

	// vld := ValidateQueryStr(c, rules)
	// if vld != nil {
	// 	return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	// }

	// result, err := models.FindAllUsers(page, rp, &models.UserFilterable{name, email})
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	// return c.JSON(http.StatusOK, result)
	return errors.New("")
}

func GetBlogById(c echo.Context) error {
	// id, err := strconv.Atoi(c.QueryParam("id"))

	// defer c.Request().Body.Close()

	// rules := govalidator.MapData{
	// 	"id": []string{"numeric"},
	// }

	// vld := ValidateQueryStr(c, rules)
	// if vld != nil {
	// 	return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	// }

	// result, err := models.FindUserByID(id)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	// return c.JSON(http.StatusOK, result)
	return errors.New("")
}

func AddBlog(c echo.Context) error {
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
	return errors.New("")
}

func EditBlog(c echo.Context) error {
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

func DeleteBlog(c echo.Context) error {
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

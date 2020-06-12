package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/NavenduDuari/go-echo-blog/src/db/otp"
	otpModel "github.com/NavenduDuari/go-echo-blog/src/db/otp/model"
	"github.com/NavenduDuari/go-echo-blog/src/db/user"
	userModel "github.com/NavenduDuari/go-echo-blog/src/db/user/model"
	"github.com/labstack/echo"
)

func SignUp(c echo.Context) error {
	newUser := new(userModel.User)
	if err := c.Bind(newUser); err != nil {
		return err
	}

	//sns action
	snsTopicArn, err := otp.CreateTopic(newUser.Email)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusForbidden, "SignUp Faild")
	}
	if err = otp.SubscribeTopic(otpModel.SubscribeSnsTopicInput{
		Endpoint: newUser.Email,
		Protocol: "email",
		TopicArn: snsTopicArn,
	}); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusForbidden, "SignUp Faild")
	}

	newUser.SnsTopicArn = snsTopicArn
	//db action
	if err := user.InsertUser(newUser); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusForbidden, "SignUp Faild")
	}

	return c.JSON(http.StatusOK, "SignUp successful")
}

func EditUser(c echo.Context) error {
	return errors.New("")
}

func DeleteUser(c echo.Context) error {
	return errors.New("")
}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/NavenduDuari/go-echo-blog/src/api/model"
	"github.com/NavenduDuari/go-echo-blog/src/common/utils"
	"github.com/NavenduDuari/go-echo-blog/src/db/auth"
	authModel "github.com/NavenduDuari/go-echo-blog/src/db/auth/model"
	"github.com/NavenduDuari/go-echo-blog/src/db/user"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	userForLogin := new(model.UserForLogin)
	if err := c.Bind(userForLogin); err != nil {
		return err
	}
	userData, err := user.FetchUserByUserId(userForLogin.UserId)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Failed Login")
	}

	if userData.Password == userForLogin.Password {
		token, err := utils.GetJwtLoginToken(model.UserJWT{
			Id:    userData.Id,
			Phone: userData.Phone,
			Email: userData.Email,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Failed Login")
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Login successful!",
			"token":   token,
		})
	}

	return c.String(http.StatusUnauthorized, "Login failed! Password didn't match")
}

func SendOtp(c echo.Context) error {
	userForLogin := new(model.UserForLogin)
	if err := c.Bind(userForLogin); err != nil {
		return err
	}
	userData, err := user.FetchUserByUserId(userForLogin.UserId)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Login failed! Unable to send OTP")
	}

	if err = auth.Publish(authModel.PublishSnsTopicInput{
		Message: authModel.OtpAuth{
			UserId:    userData.Id,
			OTP:       utils.GenerateRandomNumericText(4),
			ExpiresAt: utils.Int64ToStr(utils.GetCurrentTimeStampSecond() + 5*60),
		},
		TopicArn: userData.SnsTopicArn,
	}); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Login failed! Unable to send OTP")
	}

	return c.String(http.StatusOK, "OTP sent")
}

func VerifyOtp(c echo.Context) error {
	userForLogin := new(model.UserForLogin)
	if err := c.Bind(userForLogin); err != nil {
		return err
	}
	userData, err := user.FetchUserByUserId(userForLogin.UserId)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Failed Login")
	}
	otpAuth, err := auth.FetchOtp(userData.Id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Failed Login")
	}

	if utils.StrToInt64(otpAuth.ExpiresAt) > utils.GetCurrentTimeStampSecond() && userForLogin.OTP == otpAuth.OTP {
		token, err := utils.GetJwtLoginToken(model.UserJWT{
			Id:    userData.Id,
			Phone: userData.Phone,
			Email: userData.Email,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Failed Login")
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Login successful!",
			"token":   token,
		})
	}

	return c.String(http.StatusUnauthorized, "Login failed! Unable to verify OTP")
}

func GetQr(c echo.Context) error {
	userForLogin := new(model.UserForLogin)
	if err := c.Bind(userForLogin); err != nil {
		return err
	}
	userData, err := user.FetchUserByUserId(userForLogin.UserId)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Login failed! Unable to get QR")
	}

	qrBase64, err := auth.GenerateQr(authModel.QrAuth{
		UserId:    userData.Id,
		QrCode:    utils.GenerateRandomReadableText(10),
		ExpiresAt: utils.Int64ToStr(utils.GetCurrentTimeStampSecond() + 5*60),
	})
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Login failed! Unable to get QR")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message":  "Got QR",
		"qrBAse64": qrBase64,
	})
}

func VerifyQr(c echo.Context) error {
	userForLogin := new(model.UserForLogin)
	if err := c.Bind(userForLogin); err != nil {
		return err
	}

	userData, err := user.FetchUserByUserId(userForLogin.UserId)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Login failed! Unable to verify QR")
	}
	qrAuth, err := auth.FetchQr(userData.Id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Login failed! Unable to verify QR")
	}

	if utils.StrToInt64(qrAuth.ExpiresAt) > utils.GetCurrentTimeStampSecond() && userForLogin.QrCode == qrAuth.QrCode {
		token, err := utils.GetJwtLoginToken(model.UserJWT{
			Id:    userData.Id,
			Phone: userData.Phone,
			Email: userData.Email,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Login failed! Unable to verify QR")
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Login successful!",
			"token":   token,
		})
	}

	return c.String(http.StatusUnauthorized, "Login failed! Unable to verify QR")
}

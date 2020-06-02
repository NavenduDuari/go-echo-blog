package middlewares

import (
	"net/http"
	"strings"

	"github.com/NavenduDuari/go-echo-blog/src/common/utils"
	"github.com/labstack/echo"
)

func SetJwtMiddlewares(e *echo.Group) {
	e.Use(verifyJwt)
}

func verifyJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		splitArr := strings.Split(authToken, " ")
		jwtToken := splitArr[1]

		_, err := utils.ValidateJwtLoginToken(jwtToken)
		if err != nil {
			return c.String(http.StatusUnauthorized, "User not authorized! Please Login")
		}
		return next(c)
	}
}

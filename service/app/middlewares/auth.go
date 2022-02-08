package middlewares

import (
	"backend/app/repositories"
	"backend/app/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func AuthMiddleWare(jwtService *services.Jwt, userRepository *repositories.UserRepository) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")

			subs := strings.Split(auth, " ")
			if len(subs) < 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "your client does not have permission to the requested")
			}

			userId, err := jwtService.ParseToken(subs[1])
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "your client does not have permission to the requested")
			}

			user, err := userRepository.GetById(userId)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "your client does not have permission to the requested")
			}

			c.Set("auth", user)

			return next(c)
		}
	}
}

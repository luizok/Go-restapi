package auth

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var currentToken string

func Login(c echo.Context) error {

	u := new(UserCredentials)
	if err := c.Bind(u); err != nil {
		return err
	}

	if u.Username == "admin" && u.Password == "admin" {
		id := uuid.New()
		currentToken = id.String()
		return c.JSON(http.StatusOK, echo.Map{
			"token": currentToken,
		})
	}

	return echo.ErrUnauthorized
}

func KeyAuthHandler(key string, c echo.Context) (bool, error) {
	return key == currentToken, nil
}

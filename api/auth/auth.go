package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type JwtWithScopeClaims struct {
	Scopes []string `json:"scopes"`
	jwt.RegisteredClaims
}
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var MyJWTSecret string = "foobar"

func LoginJWT(c echo.Context) error {

	u := new(UserCredentials)
	if err := c.Bind(u); err != nil {
		return echo.ErrBadRequest
	}

	if u.Username != "admin" || u.Password != "admin" {
		return echo.ErrUnauthorized
	}

	now := time.Now()
	claims := JwtWithScopeClaims{
		[]string{UsersReadOnly, UsersReadWrite},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Minute * 30)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "myrestapi",
			Subject:   "admin",
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(MyJWTSecret))

	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenString,
	})
}

func CheckScopes(requiredScopes []string) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			token, err := jwt.ParseWithClaims(
				token.Raw,
				&JwtWithScopeClaims{},
				func(token *jwt.Token) (interface{}, error) { return []byte(MyJWTSecret), nil })

			if err != nil {
				return echo.ErrUnauthorized
			}

			claims := token.Claims.(*JwtWithScopeClaims)

			fmt.Println(claims.Scopes)
			for _, v := range claims.Scopes {
				for _, r := range requiredScopes {
					if v == r {
						return next(c)
					}
				}
			}

			return echo.ErrForbidden
		}
	}
}

package main

import (
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	auth "github.com/luizok/myrestapi/api/auth"
	models "github.com/luizok/myrestapi/api/models"
)

func main() {

	e := echo.New()

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "${time_rfc3339} ${status} ${method} ${uri} ${status} from ${remote_ip} ${latency_human}\n",
	// }))
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Hello, World!",
		})
	})

	e.POST("/login", auth.LoginJWT)

	api := e.Group(
		"/api/v1",
		echojwt.JWT([]byte(auth.MyJWTSecret)),
	)
	models.AttachUsersRoutes(api)

	e.Logger.Fatal(e.Start(":1323"))
}

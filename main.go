package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	auth "github.com/luizok/myrestapi/api/auth"
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

	e.POST("/login", auth.Login)

	e.Group("/api/v1", middleware.KeyAuth(auth.KeyAuthHandler))

	e.Logger.Fatal(e.Start(":1323"))
}

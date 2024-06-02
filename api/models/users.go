package models

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	PostsCount uint   `json:"posts_count"`
	Followers  []User `json:"followers"`
}

var allUsers map[int]User

func generateFakeUsers() map[int]User {
	users := make(map[int]User)
	for i := 1; i <= 50; i++ {
		users[i] = User{
			ID:         i,
			Name:       "User " + strconv.Itoa(i),
			Email:      "user" + strconv.Itoa(i) + "@example.com",
			PostsCount: uint(i * 10),
			Followers: []User{
				{
					ID:         i + 1,
					Name:       "Follower " + strconv.Itoa(i+1),
					Email:      "follower" + strconv.Itoa(i+1) + "@example.com",
					PostsCount: uint((i + 1) * 10),
					Followers:  nil,
				},
				{
					ID:         i + 2,
					Name:       "Follower " + strconv.Itoa(i+2),
					Email:      "follower" + strconv.Itoa(i+2) + "@example.com",
					PostsCount: uint((i + 1) * 10),
					Followers:  nil,
				},
			},
		}
	}

	return users
}

func AttachUsersRoutes(g *echo.Group) {

	allUsers = generateFakeUsers()
	users := g.Group("/users", middleware.RemoveTrailingSlash())

	users.GET("", getUsers)
	users.GET("/:id", getUser)
	users.POST("", createUser)
	users.PUT("/:id", updateUser)
	users.DELETE("/:id", deleteUser)
}

func getUsers(c echo.Context) error {
	return c.JSON(200, allUsers)
}

func getUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": "Bad Request. Invalid ID"})
	}

	if !slices.Contains(maps.Keys(allUsers), id) {
		return c.JSON(404, echo.Map{"error": "Not Found"})

	}

	user := allUsers[id]
	return c.JSON(200, user)
}

func createUser(c echo.Context) error {
	return c.JSON(200, "createUser")
}

func updateUser(c echo.Context) error {
	return c.JSON(200, "updateUser")
}

func deleteUser(c echo.Context) error {
	return c.JSON(200, "deleteUser")
}

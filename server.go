package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("hello world", rand.Intn(10))

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello3333, World!")
	})
	// e.POST("/users", saveUser)
	e.GET("/users/:id", getUser) //path parameter
	e.GET("/show", show)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

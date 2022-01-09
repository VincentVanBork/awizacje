package controllers

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	repository *sql.DB
}

func (u *User) Hello(c echo.Context) error {
	fmt.Println(u.repository)
	fmt.Println(u)

	return c.String(http.StatusOK, "Hello, World!")
}

func (u *User) GetOne(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(u.repository)
	fmt.Println(u)

	return c.String(http.StatusOK, id)
}

func (u *User) Register(c echo.Context) error {
	fmt.Println(u.repository)
	fmt.Println(u)
	return c.String(http.StatusOK, "REGISTERED")
}

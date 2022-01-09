package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

func SetApiUrls(g *echo.Group, db *sql.DB) {
	user := User{repository: db}
	g.GET("/", user.Hello)
	g.GET("/users/:id", user.GetOne)
}

func SetUserUrls(e *echo.Echo, db *sql.DB) {
	user := User{repository: db}
	e.POST("/register", user.Register)
}

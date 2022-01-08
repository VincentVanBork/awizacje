package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"users/controllers"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", controllers.Hello)

	port, isSet := os.LookupEnv("USER_SERVICE_PORT")
	fmt.Println(port)
	if isSet {
		e.Logger.Fatal(e.Start(":" + port))
	} else {
		e.Logger.Fatal(e.Start(":54301"))
	}

}

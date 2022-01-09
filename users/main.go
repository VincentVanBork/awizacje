package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"os"
	"users/controllers"
	"users/models"
	"users/users_middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	// db and migrations
	dbConnect := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_ADDRESS"),
		os.Getenv("POSTGRES_DB"))
	fmt.Println(dbConnect)
	db, sqlErr := sql.Open("postgres", dbConnect)
	if sqlErr != nil {
		e.Logger.Fatal(sqlErr)
	}
	migrateErr := models.Migrate(db)
	if migrateErr != nil && migrateErr != migrate.ErrNoChange {
		e.Logger.Fatal(migrateErr)
	}
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	controllers.SetUserUrls(e, db)

	g := e.Group("/api")
	middAuth := users_middleware.DBMiddleware{Repository: db}
	g.Use(middleware.BasicAuth(middAuth.Login))
	controllers.SetApiUrls(g, db)

	port, isSet := os.LookupEnv("USER_SERVICE_PORT")
	if isSet {
		e.Logger.Fatal(e.Start(":" + port))
	} else {
		e.Logger.Fatal(e.Start(":54301"))
	}

}

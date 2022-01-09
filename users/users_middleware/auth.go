package users_middleware

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type DBMiddleware struct {
	Repository *sql.DB
}

func (m *DBMiddleware) Login(username, password string, c echo.Context) (bool, error) {
	return m.AuthUser(username, password)
}

// TODO: move this repository to some structs and use composition
func (m *DBMiddleware) AuthUser(email string, password string) (bool, error) {
	rows, queryErr := m.Repository.Query("SELECT * FROM users WHERE email = $e", email)
	if queryErr != nil {
		return false, queryErr
	}
	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}(rows)

	for rows.Next() {
		var authPassword string
		err := rows.Scan(&authPassword)
		if err != nil {
			log.Fatal(err)
		}
		authErr := bcrypt.CompareHashAndPassword([]byte(authPassword), []byte(password))
		if authErr != nil {
			return false, nil
		}
	}
	rowsErr := rows.Err()
	if rowsErr != nil {
		return false, rowsErr
	}
	return true, nil
}

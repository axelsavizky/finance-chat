package database

import (
	"database/sql"
	"errors"
	"finance-chat/src/configuration"
	financeChatErrors "finance-chat/src/domain/errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	searchPasswordByUsernameSQL = "SELECT password FROM users WHERE username = ?;"
	searchUserByUsernameSQL     = "SELECT id FROM users WHERE username = ?;"
	insertUserSQL               = "INSERT INTO users (username, password) VALUES (?, ?);"
)

func Login(username string) (hashedPassword string, err error) {
	db, err := connectToDatabase()
	if err != nil {
		return "", err
	}
	defer db.Close()

	row := db.QueryRow(searchPasswordByUsernameSQL, username,)

	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", financeChatErrors.NewErrNotFound("username or password is wrong")
		}
		return "", fmt.Errorf("error executing query for username: %s", username)
	}

	return hashedPassword, nil
}

func Signup(username, password string) error {
	db, err := connectToDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	row := db.QueryRow(searchUserByUsernameSQL, username)
	var id int64
	err = row.Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return errors.New("error preparing statement")
	}

	if err == nil {
		// Err must be sql.ErrNotRows
		return financeChatErrors.NewErrBadRequest("user already exists")
	}

	// err == sql.ErrNotRows so we continue
	_, err = db.Exec(insertUserSQL, username, password)
	if err != nil {
		return errors.New("error inserting in the database")
	}

	return nil
}

func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open(configuration.DatabaseDriverName, fmt.Sprintf("%s:%s@/%s",configuration.DatabaseUsername, configuration.DatabasePassword, configuration.DatabaseSchemaName))
	if err != nil {
		return nil, errors.New("error connecting to the database")
	}

	return db, nil
}

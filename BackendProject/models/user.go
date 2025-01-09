package models

import (
	"database/sql"
	"fmt"

	"github.com/backend/db"
	"github.com/backend/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() (User, error) {
	query := `
		INSERT INTO users (email, password) VALUES (?, ?)
	`

	statement, prepareErr := db.DB.Prepare(query)

	fmt.Println("created user: ==> ", u)

	if prepareErr != nil {
		fmt.Println("[1]")
		return User{}, prepareErr
	}

	finalPassword, hashError := utils.HashPassword(u.Password)

	if hashError != nil {
		fmt.Println("[2]")
		return User{}, hashError
	}

	result, execErr := statement.Exec(u.Email, finalPassword)

	defer statement.Close()

	if execErr != nil {
		fmt.Println("[3]")
		return User{}, execErr
	}

	id, idErr := result.LastInsertId()
	u.ID = id

	return u, idErr
}

func (u *User) ValidateCredentials() error {
	query := `
		SELECT id, password FROM users WHERE email = ?
	`

	var retrievedPassword string

	row := db.DB.QueryRow(query, u.Email)
	errorNoRows := row.Scan(&u.ID, &retrievedPassword)

	if errorNoRows != nil {
		return fmt.Errorf("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return fmt.Errorf("invalid credentials")
	}

	return nil

}

func (u User) DeleteUserByID() (sql.Result, error) {
	query := "DELETE FROM users WHERE id = ?"

	stmt, paraErr := db.DB.Prepare(query)

	if paraErr != nil {
		return nil, paraErr
	}

	result, execErr := stmt.Exec(u.ID)

	defer stmt.Close()

	return result, execErr
}

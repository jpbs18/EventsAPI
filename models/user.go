package models

import (
	"Events-API/db"
	"Events-API/utils"
	"errors"
)

type User struct {
	ID       int64 		`json:"id"`
	Email    string   `json:"email" binding:"required"`
	Password string   `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES(?, ?)`

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := db.DB.Exec(query, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId

	return err 
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("Invalid credentials.")
	}

	passwordIsValid := utils.CheckPasswordHash(retrievedPassword, u.Password)

	if !passwordIsValid {
		return errors.New("Invalid credentials.")
	}

	return nil
}
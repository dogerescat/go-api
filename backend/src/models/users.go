package models

import (
	"go-api/database"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Create() {
	cmd := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`
	stmt, err := database.Db.Prepare(cmd)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(u.Name, u.Email, u.Password)
	if err != nil {
		panic(err)
	}
}

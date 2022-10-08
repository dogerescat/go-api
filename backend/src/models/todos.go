package models

import (
	"go-api/database"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"id_done"`
	UserId int    `json:"user_id"`
}

func (t *Todo) Create() {
	cmd := `INSERT INTO todos (title, is_done, user_id) VALUES (?, ?, ?)`
	stmt, err := database.Db.Prepare(cmd)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(t.Title, t.IsDone, t.UserId)
	if err != nil {
		panic(err)
	}
}

func GetAll(user_id int) ([]Todo, error) {
	cmd := `SELECT * FROM todos WHERE user_id = ?`
	rows, err := database.Db.Query(cmd, user_id)
	defer rows.Close()
	var todos []Todo
	if err != nil {
		return todos, err
	}
	for rows.Next() {
		todo := Todo{}
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.IsDone, &todo.UserId); err != nil {
			panic(err)
		}
		todos = append(todos, todo)
	}
	err = rows.Err()

	return todos, err
}

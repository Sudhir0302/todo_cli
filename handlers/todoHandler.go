package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Sudhir0302/todo_cli/config"
	"github.com/Sudhir0302/todo_cli/models"
)

type TodoHandler struct {
	DB *sql.DB
}

func Init() {
	handler := TodoHandler{DB: config.DB}

	_, err := handler.DB.Exec(`CREATE TABLE IF NOT EXISTS todo(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task VARCHAR(200),
		status BOOLEAN default false
		);`)

	if err != nil {
		fmt.Println(err)
	}
}

func (todo *TodoHandler) View() {
	row, err := todo.DB.Query("SELECT * FROM todo")

	if err != nil {
		log.Println(err)
	}

	for row.Next() {
		var t models.Todo
		row.Scan(&t.Id, &t.Task, &t.Status)
		fmt.Println(t)
	}
}

func (todo *TodoHandler) Add(task string) string {
	stmt, _ := todo.DB.Prepare("INSERT INTO todo (task,status) values(?,false)")

	_, err := stmt.Exec(task)

	if err != nil {
		log.Println("failed")
		return "failed"
	}
	return "success"
}

func (todo *TodoHandler) Update(task string) string {
	stmt, _ := todo.DB.Prepare("UPDATE todo SET status=true WHERE task=?")
	res, err := stmt.Exec(task)

	if err != nil {
		// log.Println("failed")
		return "failed"
	}

	c, _ := res.RowsAffected()
	if c == 0 {
		return "invalid task"
	}
	return "success"
}

func (todo *TodoHandler) Delete(task string) string {
	stmt, _ := todo.DB.Prepare("DELETE FROM todo WHERE task=?")
	res, err := stmt.Exec(task)
	if err != nil {
		return "failed"
	}

	c, _ := res.RowsAffected()
	if c == 0 {
		return "invalid task"
	}

	return "success"
}

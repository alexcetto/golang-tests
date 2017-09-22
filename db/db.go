package db

import (
	"../types"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

var err error
var database *sql.DB

func init() {
	defer database.Close()
	database, err = sql.Open("sqlite3", "./identifier.sqlite")
	if err == nil {
		log.Println(err)
	}
}

func GetTasks() types.Context {
	var tasks []types.Task
	var context types.Context
	var TaskID int
	var TaskTitle string
	var TaskContent string
	var TaskCreated time.Time
	var getTaskSQL string

	getTaskSQL = "SELECT id, title, content, created_date FROM task;"

	rows, err := database.Query(getTaskSQL)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&TaskID, &TaskTitle, &TaskContent, &TaskCreated)
		if err != nil {
			log.Println(err)
		}
		TaskCreated = TaskCreated.Local()
		a := types.Task{Id: TaskID, Title: TaskTitle, Content: TaskContent, Created: TaskCreated.Format(time.UnixDate)[0:20]}
		tasks = append(tasks, a)
	}
	context = types.Context{Tasks: tasks}
	return context
}

package db

import (
	"database/sql"
	"github.com/alexcetto/golang-tests/types"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
	"fmt"
)

var err error
var database *sql.DB

func init() {
	database, err = sql.Open("sqlite3", "/Users/alexandrecetto/go/src/github.com/alexcetto/golang-tests/schema.db")
	if err != nil {
		log.Println("We are here now")
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

func Addtask(title, content string) error {
	query := "INSERT INTO task(title, content, created_date, last_modified_at) values(?,?, datetime(), datetime())"
		restoreSQL, err := database.Prepare(query)
		if err != nil {
			fmt.Println(err)
		}

		tx, err := database.Begin()
		_, err = tx.Stmt(restoreSQL).Exec(title, content)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
		} else {
			log.Print("Insert Successful")
			tx.Commit()
		}
		return err
}
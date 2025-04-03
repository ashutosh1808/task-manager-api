package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "tasks.db")
	if err != nil {
		log.Println("Error opening sqlite: ", err)
	}

	DB.SetMaxIdleConns(5)
	DB.SetConnMaxIdleTime(10)
	createTable()
}

func createTable() {
	createTasksTable := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		task TEXT NOT NULL,
		description TEXT NOT NULL,
		status TEXT NOT NULL,
		dueDate DATETIME
	)
	`

	_, err := DB.Exec(createTasksTable)
	if err != nil {
		log.Println(err)
	}
}

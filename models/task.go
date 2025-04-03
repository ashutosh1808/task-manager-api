package models

import (
	"time"

	"example.com/investment-calculator/practice/task-manager-api/db"
)

type Task struct {
	ID          int64
	Task        string    `binding:"required"`
	Description string    `binding:"required"`
	Status      string    `binding:"required"`
	DueDate     time.Time `binding:"required"`
}

func (t Task) Save() error {
	query := `INSERT INTO tasks(task,description,status,dueDate) values(?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(t.Task, t.Description, t.Status, t.DueDate)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = id
	return nil
}

func GetTasks() ([]Task, error) {
	var tasks []Task

	query := `SELECT id,task,description,status,dueDate from tasks`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Task, &task.Description, &task.Status, &task.DueDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskById(id int64) (*Task, error) {
	query := `SELECT * from tasks where id=?`
	row := db.DB.QueryRow(query, id)

	var task Task
	err := row.Scan(&task.ID, &task.Task, &task.Description, &task.Status, &task.DueDate)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (t Task) UpdateTask() error {
	query := `UPDATE tasks SET task=?, description=?, status=?, dueDate = ? where id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Task, t.Description, t.Status, t.DueDate, t.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t Task) DeleteTask() error {
	query := `DELETE from tasks where id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(t.ID)
	return err
}

package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type TaskJson struct {
	Task_id          string `json:"task_id"`
	Task_name        string `json:"task_name"`
	Task_description string `json:"task_description"`
	User_id          string `json:"user_id"`
	Created_at       string `json:"created_at"`
}

func CreateTasksTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS tasks (
    task_id VARCHAR(65) NOT NULL,
    task_name VARCHAR(25) NOT NULL,
    task_desc VARCHAR(255),
    user_id VARCHAR(129) REFERENCES "User"(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW()
);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("couldn't create users table:%v", err)
	}
}

//

func InsertInTasks(w http.ResponseWriter, db *sql.DB, task *TaskJson) {
	if task.Task_name == "" {
		AnswerWithError(w, 400, fmt.Sprintf("task_name field is empty"))
		return
	}
	if task.User_id == "" {
		AnswerWithError(w, 400, fmt.Sprintf("user_id field is empty"))
		return
	}
	taskId, err := uuid.NewUUID()
	query := `INSERT INTO tasks(task_id,task_name,task_desc,user_id)
	VALUES ($1,$2,$3,$4) RETURNING *`
	if err != nil {
		AnswerWithError(w, 400, fmt.Sprintf("couldn't create uuid for task table: %v", err))
		return
	}
	var (
		task_id    string
		task_name  string
		task_desc  string
		user_id    string
		created_at string
	)
	err = db.QueryRow(query, taskId.String(), task.Task_name, task.Task_description, task.User_id).Scan(&task_id, &task_name, &task_desc, &user_id, &created_at)
	if err != nil {
		AnswerWithError(w, 400, fmt.Sprintf("couldn't insert Task data in database: %v", err))
		return
	}
	created_at = parseTime(created_at)
	AnswerWithJson(w, 200, &TaskJson{
		Task_id:          task_id,
		Task_name:        task_name,
		Task_description: task_desc,
		User_id:          user_id,
		Created_at:       created_at,
	})
}

type Task struct {
	TaskID    string `json:"task_id"`
	TaskName  string `json:"task_name"`
	TaskDesc  string `json:"task_description"`
	CreatedAt string `json:"created_at"`
}

func GetAllToDo(w http.ResponseWriter, db *sql.DB, TaskStruct *TaskJson) {
	if TaskStruct.User_id == "" {
		AnswerWithError(w, 400, fmt.Sprintf("no user_id inserted"))
		return
	}
	query := `SELECT task_id,task_name,task_desc,to_char(created_at, 'YYYY-MM-DD HH24:MI:SS') FROM tasks                                                                     
    WHERE user_id=$1 ORDER BY created_at DESC `
	rows, err := db.Query(query, TaskStruct.User_id)
	if err != nil {
		AnswerWithError(w, 400, fmt.Sprintf("wrong id key or ivalid format: %v", err))
		return
	}
	var (
		task_id    string
		task_name  string
		task_desc  string
		created_at string
	)
	defer rows.Close()
	tasks := []Task{}
	for rows.Next() {
		err := rows.Scan(&task_id, &task_name, &task_desc, &created_at)
		if err != nil {
			AnswerWithError(w, 500, fmt.Sprintf("error scanning row: %v", err))
			return
		}
		tasks = append(tasks, Task{task_id, task_name, task_desc, created_at})
	}
	if err := rows.Err(); err != nil {
		AnswerWithError(w, 500, fmt.Sprintf("error iterating over rows: %v", err))
		return
	}
	AnswerWithJson(w, 200, tasks)
}

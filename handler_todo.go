package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type dbFunctionTodo func(w http.ResponseWriter, db *sql.DB, TaskStruct *TaskJson)

func (apiCfg *MyDB) HandlerCreateTodo(dbFunc dbFunctionTodo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body) //инициализирует декодер
		params := &TaskJson{}
		err := decoder.Decode(&params) //читает боди реквеста
		if err != nil {
			AnswerWithError(w, 400, fmt.Sprintf("Error parsing json:: %v", err))
			return
		}
		dbFunc(w, apiCfg.Db, params)
	}
}

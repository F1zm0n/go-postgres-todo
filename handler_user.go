package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type MyDB struct {
	Db *sql.DB
}
type dbFunction func(w http.ResponseWriter, db *sql.DB, user *User)

func (apiCfg *MyDB) HandlerCreateUser(dbFunc dbFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body) //инициализирует декодер
		params := &User{}
		err := decoder.Decode(&params) //читает боди реквеста
		if err != nil {
			AnswerWithError(w, 400, fmt.Sprintf("Error parsing json:: %v", err))
			return
		}
		dbFunc(w, apiCfg.Db, params)
	}
}

//YmltYmlt/Z2hnaGcwMzA1QGdtYWlsLmNvbQ==

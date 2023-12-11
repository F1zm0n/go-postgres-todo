package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	Db "github.com/F1zm0n/auth.git/myDb"
	"net/http"
)

type MyDB struct {
	Db *sql.DB
}
type dbFunction func(w http.ResponseWriter, db *sql.DB, user Db.User)

func (apiCfg *MyDB) HandlerCreateUser(dbFunc dbFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body) //инициализирует декодер
		params := Db.User{}
		err := decoder.Decode(&params) //читает боди реквеста
		if err != nil {
			AnswerWithError(w, 400, fmt.Sprintf("Error parsing json:: %v", err))
			return
		}
		dbFunc(w, apiCfg.Db, params)
	}
}

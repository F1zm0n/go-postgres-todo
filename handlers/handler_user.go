package handlers

import (
	"encoding/json"
	"fmt"
	Db "github.com/F1zm0n/auth.git/myDb"
	"net/http"
)

func HandlerCreateUser(w http.ResponseWriter, r http.Request) {
	decoder := json.NewDecoder(r.Body) //инициализирует декодер
	params := Db.User{}
	err := decoder.Decode(&params) //читает боди реквеста
	if err != nil {
		answerWithError(w, 400, fmt.Sprintf("Error parsing json:: %v", err))
		return
	}
	Db.InsertInUserTable(DataB)
}

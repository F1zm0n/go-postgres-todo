package main

import (
	"fmt"
	"net/http"
)

// middlewareAuth должен принимать authedHandler и возвращать http.HandlerFunc с правильной сигнатурой
// что бы она подходила сигнатуре
func (apiCfg *MyDB) middlewareAuth(dbFunc dbFunction) http.HandlerFunc {
	//возвращаем хэндлер с правильной сигнатурой
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := GetIdKey(r.Header)
		if err != nil {
			AnswerWithError(w, 400, fmt.Sprintf("couldn't get id from header:", err))
			return
		}
		idStruct := &User{Id: id}
		dbFunc(w, apiCfg.Db, idStruct)
	}
}

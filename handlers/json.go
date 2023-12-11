package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func AnswerWithJson(w http.ResponseWriter, code int, parser interface{}) {
	jsoned, err := json.Marshal(parser)
	if err != nil {
		log.Printf("couldn't marshall json %v, error: %v", parser, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(jsoned)
	if err != nil {
		log.Printf("couldn't write json: %v, error: %v", jsoned, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
}
func AnswerWithError(w http.ResponseWriter, code int, msg string) { //функция выведения ошибок выше 500
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	AnswerWithJson(w, code, errResponse{
		Error: msg,
	})
}

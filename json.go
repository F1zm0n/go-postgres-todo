package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func AnswerWithJson(w http.ResponseWriter, code int, payload interface{}) {
	// Преобразование структуры данных (payload) в формат JSON.
	dat, err := json.Marshal(payload)
	// Если произошла ошибка при маршалинге, логируем сообщение и отправляем HTTP-ответ с кодом 500 (Internal Server Error).
	if err != nil {
		log.Printf("Failed to marshall JSON response %v", payload)
		w.WriteHeader(500)
		return
	}
	// Установка заголовка Content-type на "application/json".
	w.Header().Add("Content-type", "application/json")
	// Установка кода HTTP-ответа
	w.WriteHeader(code)
	// Запись данных в тело HTTP-ответа.
	_, err = w.Write(dat)
	if err != nil {
		log.Printf("Failed to write body %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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

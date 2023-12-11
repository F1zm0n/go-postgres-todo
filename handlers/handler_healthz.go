package handlers

import (
	"net/http"
)

func HandlerHealth(w http.ResponseWriter, r *http.Request) {
	AnswerWithJson(w, 200, struct{}{})
}

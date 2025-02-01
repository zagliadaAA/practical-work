package controller

import (
	"encoding/json"
	"net/http"
)

// RespondJSON отправляет JSON-ответ клиенту с указанным кодом статуса.
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

// RespondValidationError отправляет JSON-ответ с ошибкой.
func RespondValidationError(w http.ResponseWriter, v *ValidationError) {
	RespondJSON(w, http.StatusBadRequest, v)
}

// RespondInternalServerError отправляет JSON-ответ с сообщением об ошибке сервера.
func RespondInternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

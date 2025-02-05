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

func RespondStatusBadRequestError(w http.ResponseWriter, s *StatusBadRequestError) {
	RespondJSON(w, http.StatusBadRequest, s)
}

// RespondInternalServerError отправляет JSON-ответ с сообщением об ошибке сервера.
func RespondInternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func RespondNotFoundError(w http.ResponseWriter) {
	RespondJSON(w, http.StatusNotFound, nil)
}

// DecodeRequest считывает тело запроса и помещает его в структуру
func DecodeRequest(w http.ResponseWriter, r *http.Request, req interface{}) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		RespondValidationError(w, NewValidationError("decode", "error reading json"))

		return err
	}

	return nil
}

// EncodeResponse записывает данные из структуры в json и возвращает в качестве ответа
func EncodeResponse(w http.ResponseWriter, response interface{}) error {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := encoder.Encode(response)
	if err != nil {
		RespondValidationError(w, NewValidationError("encode", "error writing json encoding"))

		return err
	}

	return nil
}

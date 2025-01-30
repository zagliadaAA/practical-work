package client_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"project2/internal/usecase/client_usecase"
)

type createClientReq struct {
	Name        string    `json:"name"`
	BirthDate   time.Time `json:"birth_date"`
	PhoneNumber string    `json:"phone_number"`
}

type createClientResp struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	BDate       time.Time `json:"birth_date"`
	PhoneNumber string    `json:"phone_number"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	// парсим json в нашу структуру
	// валидируем тело запроса или парамерты
	// вызываем юзкейс
	// обрабатываем ошибки если есть
	// возвращаем результат

	decoder := json.NewDecoder(r.Body)
	var req createClientReq
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err = validateCreateClientReq(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	client, err := c.clientUseCase.Create(client_usecase.CreateClientReq{
		Name:        req.Name,
		BDate:       req.BirthDate,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = encoder.Encode(createClientResp{
		ID:          client.ID,
		Name:        client.Name,
		BDate:       client.BDate,
		PhoneNumber: client.PhoneNumber,
		UpdatedAt:   client.UpdatedAt,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func validateCreateClientReq(r createClientReq) error {
	if r.Name == "" {
		return errors.New("name is required")
	}

	if r.PhoneNumber == "" || len(r.PhoneNumber) > 11 {
		return errors.New("phone_number is required")
	}

	if r.BirthDate.IsZero() {
		return errors.New("birth_date is required")
	}

	return nil
}

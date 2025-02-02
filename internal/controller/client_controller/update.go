package client_controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"medicalCenter/internal/controller"
	"medicalCenter/internal/usecase/client_usecase"
)

type updateClientReq struct {
	Name        string `json:"name"`
	BirthDate   string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
}

type updateClientResp struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	BDate       time.Time `json:"birth_date"`
	PhoneNumber string    `json:"phone_number"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Path[len("/clients/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid ID client", http.StatusBadRequest)

		return
	}

	var req updateClientReq
	if err := c.DecodeRequest(w, r, &req); err != nil {
		return
	}

	validationError := validateUpdateClientReq(req)
	if validationError != nil {
		controller.RespondValidationError(w, validationError)

		return
	}

	//преобразование даты из строки в time.time
	birthDate, err := time.Parse(time.DateOnly, req.BirthDate)
	if err != nil {
		controller.RespondValidationError(w, controller.NewValidationError("birth_date", "time format (2000-11-11) not allowed"))

		return
	}

	client, err := c.clientUseCase.Update(client_usecase.UpdateClientReq{
		ID:          id,
		Name:        req.Name,
		BDate:       birthDate,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = encoder.Encode(updateClientResp{
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

func validateUpdateClientReq(r updateClientReq) *controller.ValidationError {
	if r.Name == "" || len(r.Name) > 120 {
		return controller.NewValidationError("name", "name not null, lenght no more then 120")
	}

	if r.PhoneNumber == "" || len(r.PhoneNumber) > 11 {
		return controller.NewValidationError("phone_number", "phone number not null, lenght no more then 11")
	}

	if r.BirthDate == "" {
		return controller.NewValidationError("birth_date", "birth_date not null")
	}

	return nil
}

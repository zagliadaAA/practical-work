package client_controller

import (
	"net/http"
	"strconv"
	"time"

	"medicalCenter/internal/controller"
)

type getClientByIDResp struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	BDate       time.Time `json:"birth_date"`
	PhoneNumber string    `json:"phone_number"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (c *Controller) GetClientByID(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Path[len("/clients/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		controller.RespondStatusBadRequestError(w, controller.NewStatusBadRequestError("failed converted to type int"))

		return
	}

	client, err := c.clientUseCase.GetClientByID(id)
	if err != nil {
		controller.RespondStatusBadRequestError(w, controller.NewStatusBadRequestError("failed to get client"))

		return
	}

	if err = controller.EncodeResponse(w, mapClientToResponseForGetClientByID(client)); err != nil {
		return
	}
}

package client_controller

import (
	"net/http"
	"strconv"

	"medicalCenter/internal/controller"
)

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		controller.RespondValidationError(w, controller.NewValidationError("id", "invalid id"))

		return
	}

	err = c.clientUseCase.Delete(id)
	if err != nil {
		controller.RespondStatusBadRequestError(w, controller.NewStatusBadRequestError("failed to delete client"))

		return
	}

	w.WriteHeader(http.StatusOK)
}

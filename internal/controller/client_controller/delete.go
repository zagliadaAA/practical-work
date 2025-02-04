package client_controller

import (
	"net/http"
	"strconv"

	"medicalCenter/internal/controller"
)

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Path[len("/clients/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		controller.RespondValidationError(w, controller.NewValidationError("atoi", "failed converted to type int"))

		return
	}

	err = c.clientUseCase.Delete(id)
	if err != nil {
		controller.RespondValidationError(w, controller.NewValidationError("delete", "failed to delete client"))

		return
	}

	w.WriteHeader(http.StatusOK)
}

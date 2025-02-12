package medical_report_controller

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

	err = c.medicalReportUseCase.Delete(id)
	if err != nil {
		controller.RespondStatusBadRequestError(w, controller.NewStatusBadRequestError("failed to delete medical report"))

		return
	}

	w.WriteHeader(http.StatusOK)
}

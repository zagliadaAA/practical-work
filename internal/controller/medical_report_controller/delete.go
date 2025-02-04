package medical_report_controller

import (
	"net/http"
	"strconv"

	"medicalCenter/internal/controller"
)

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Path[len("/reports/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		controller.RespondValidationError(w, controller.NewValidationError("atoi", "failed converted to type int"))

		return
	}

	err = c.medicalReportUseCase.Delete(id)
	if err != nil {
		controller.RespondValidationError(w, controller.NewValidationError("delete", "failed to delete medical report"))

		return
	}

	w.WriteHeader(http.StatusOK)
}

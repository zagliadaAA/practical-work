package medical_report_controller

import (
	"net/http"
	"strconv"
	"time"

	"medicalCenter/internal/controller"
)

type getReportByIDResp struct {
	ID         int       `json:"id"`
	DoctorName string    `json:"doctor_name"`
	Diagnosis  string    `json:"diagnosis"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IdClient   int       `json:"id_client"`
}

func (c *Controller) GetReportByID(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Path[len("/reports/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		controller.RespondStatusBadRequestError(w, controller.NewStatusBadRequestError("failed converted to type int"))

		return
	}

	report, err := c.medicalReportUseCase.GetReportByID(id)
	if err != nil {
		controller.RespondStatusBadRequestError(w, controller.NewStatusBadRequestError("failed to get medical report"))

		return
	}

	if err = controller.EncodeResponse(w, mapReportToResponseForGetReportByID(report)); err != nil {
		return
	}
}

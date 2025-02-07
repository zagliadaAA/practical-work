package medical_report_controller

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"medicalCenter/internal/adapter/postgres"
	"medicalCenter/internal/controller"
	"medicalCenter/internal/usecase/medical_report_usecase"
)

type updateMedicalReportReq struct {
	IDClient   int    `json:"id_client"`
	DoctorName string `json:"doctor_name"`
	Diagnosis  string `json:"diagnosis"`
}

type updateMedicalReportResp struct {
	ID         int       `json:"id"`
	DoctorName string    `json:"doctor_name"`
	Diagnosis  string    `json:"diagnosis"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IdClient   int       `json:"id_client"`
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		controller.RespondValidationError(w, controller.NewValidationError("id", "invalid id"))

		return
	}

	var req updateMedicalReportReq
	if err := controller.DecodeRequest(w, r, &req); err != nil {
		return
	}

	validationError := validateUpdateMedicalReportReq(&req)
	if validationError != nil {
		controller.RespondValidationError(w, validationError)

		return
	}

	medicalReport, err := c.medicalReportUseCase.Update(medical_report_usecase.UpdateReportReq{
		ID:         id,
		IDClient:   req.IDClient,
		DoctorName: req.DoctorName,
		Diagnosis:  req.Diagnosis,
	})
	if err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			controller.RespondNotFoundError(w)

			return
		}
		controller.RespondStatusBadRequestError(w, controller.NewStatusBadRequestError("failed to update medical report"))

		return
	}

	if err = controller.EncodeResponse(w, mapMedicalReportToResponseForUpdate(medicalReport)); err != nil {
		return
	}
}

func validateUpdateMedicalReportReq(r *updateMedicalReportReq) *controller.ValidationError {
	if r.IDClient == 0 {
		return controller.NewValidationError("idClient", "client id should not be 0")
	}

	if r.DoctorName == "" || len(r.DoctorName) > 120 {
		return controller.NewValidationError("DoctorName", "DoctorName not null, lenght no more then 120")
	}

	if r.Diagnosis == "" || len(r.Diagnosis) > 6 {
		return controller.NewValidationError("Diagnosis", "Diagnosis not null, lenght no more then 6")
	}

	return nil
}

package medical_report_controller

import (
	"encoding/json"
	"net/http"
	"time"

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
	var req updateMedicalReportReq
	if err := c.DecodeRequest(w, r, &req); err != nil {
		return
	}

	validationError := validateUpdateMedicalReportReq(&req)
	if validationError != nil {
		controller.RespondValidationError(w, validationError)

		return
	}

	medicalReport, err := c.medicalReportUseCase.Update(medical_report_usecase.UpdateReportReq{
		IDClient:   req.IDClient,
		DoctorName: req.DoctorName,
		Diagnosis:  req.Diagnosis,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = encoder.Encode(updateMedicalReportResp{
		ID:         medicalReport.ID,
		DoctorName: medicalReport.DoctorName,
		Diagnosis:  medicalReport.Diagnosis,
		CreatedAt:  medicalReport.CreatedAt,
		UpdatedAt:  medicalReport.UpdatedAt,
		IdClient:   medicalReport.IDClient,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

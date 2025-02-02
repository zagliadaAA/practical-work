package medical_report_controller

import (
	"encoding/json"
	"net/http"

	"medicalCenter/internal/domain"
	"medicalCenter/internal/usecase/medical_report_usecase"
)

type Controller struct {
	medicalReportUseCase medicalReportUseCase
}

type medicalReportUseCase interface {
	Create(req medical_report_usecase.CreateMedicalReportReq) (*domain.MedicalReport, error)
	Delete(id int) error
	Update(req medical_report_usecase.UpdateReportReq) (*domain.MedicalReport, error)
}

func NewController(useCase medicalReportUseCase) *Controller {
	return &Controller{
		medicalReportUseCase: useCase,
	}
}

func (c *Controller) DecodeRequest(w http.ResponseWriter, r *http.Request, req interface{}) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return err
	}

	return nil
}

package medical_report_controller

import (
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

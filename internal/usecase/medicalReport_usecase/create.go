package medicalReport_usecase

import (
	"fmt"
	"project2/internal/domain"
)

type CreateMedicalReportReq struct { //вспомогательный тип для создания без ID
	IDClient   int
	DoctorName string
	Diagnosis  string
}

func (uc *UseCase) Create(req CreateMedicalReportReq) error {
	report := domain.NewMedicalReport(req.DoctorName, req.Diagnosis)

	report.IDClient = req.IDClient

	err := uc.medRepo.Create(report)
	if err != nil {
		return fmt.Errorf("create report: %v", err)
	}

	return nil
}

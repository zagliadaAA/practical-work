package medical_report_usecase

import (
	"fmt"

	"medicalCenter/internal/domain"
)

type CreateMedicalReportReq struct { //Вспомогательный тип для создания без ID
	IDClient   int
	DoctorName string
	Diagnosis  string
}

func (uc *UseCase) Create(req CreateMedicalReportReq) (*domain.MedicalReport, error) {
	client, err := uc.clientRepo.FindByID(req.IDClient)
	if err != nil {
		return nil, fmt.Errorf("clientRepo.FindByID: %w", err)
	}

	report := domain.NewMedicalReport(req.DoctorName, req.Diagnosis, client.ID, uc.timer.Now(), uc.timer.Now())

	reportCreate, err := uc.medRepo.Create(report)
	if err != nil {
		return nil, fmt.Errorf("medRepo.Create: %w", err)
	}

	return reportCreate, nil
}

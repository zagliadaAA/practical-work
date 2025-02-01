package medical_report_usecase

import (
	"fmt"

	"medicalCenter/internal/domain"
)

type UpdateReportReq struct { //Вспомогательный тип для создания без ID
	DoctorName string
	Diagnosis  string
	IDClient   int
}

func (uc *UseCase) Update(req UpdateReportReq) (*domain.MedicalReport, error) {
	report, err := uc.medRepo.GetReportByIDClient(req.IDClient)
	if err != nil {
		return nil, fmt.Errorf("medRepo.GetReportByIDClient: %w", err)
	}

	report.DoctorName = req.DoctorName
	report.Diagnosis = req.Diagnosis
	report.UpdatedAt = uc.timer.Now()

	reportUpdate, err := uc.medRepo.Update(report)
	if err != nil {
		return nil, fmt.Errorf("medRepo.Update: %w", err)
	}

	return reportUpdate, nil
}

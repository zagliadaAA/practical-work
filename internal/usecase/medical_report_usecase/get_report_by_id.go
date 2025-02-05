package medical_report_usecase

import (
	"fmt"

	"medicalCenter/internal/domain"
)

func (uc *UseCase) GetReportByID(id int) (*domain.MedicalReport, error) {
	report, err := uc.medRepo.GetReportByID(id)
	if err != nil {
		return nil, fmt.Errorf("medRepo.GetReportByID: %w", err)
	}

	return report, nil
}

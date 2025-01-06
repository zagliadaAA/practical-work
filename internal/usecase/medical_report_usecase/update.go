package medical_report_usecase

import (
	"fmt"
)

type UpdateReportReq struct { //Вспомогательный тип для создания без ID
	DoctorName string
	Diagnosis  string
	IDClient   int
}

func (uc *UseCase) Update(req UpdateReportReq) error {
	report, err := uc.medRepo.GetReportByIDClient(req.IDClient)
	if err != nil {
		return fmt.Errorf("medRepo.GetReportByIDClient: %w", err)
	}

	report.DoctorName = req.DoctorName
	report.Diagnosis = req.Diagnosis

	if err = uc.medRepo.Update(report); err != nil {
		return fmt.Errorf("medRepo.Update: %w", err)
	}

	return nil
}

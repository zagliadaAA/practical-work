package medical_report_usecase

import (
	"fmt"
)

type UpdateReportReq struct { //вспомогательный тип для создания без ID
	DoctorName string
	Diagnosis  string
	IDClient   int
}

func (uc *UseCase) Update(req UpdateReportReq) error {
	reportID, err := uc.medRepo.GetIdReport(req.IDClient)
	if err != nil {
		return fmt.Errorf("medRepo.GetIdReport: %w", err)
	}

	report, err := uc.medRepo.FindByID(reportID)
	if err != nil {
		return fmt.Errorf("medRepo.FindByID: %w", err)
	}

	report.DoctorName = req.DoctorName
	report.Diagnosis = req.Diagnosis

	if err := uc.medRepo.Update(report); err != nil {
		return fmt.Errorf("medRepo.Update: %w", err)
	}

	return nil
}

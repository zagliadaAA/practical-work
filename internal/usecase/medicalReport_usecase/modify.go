package medicalReport_usecase

import (
	"fmt"
	"project2/internal/domain"
)

type ModifyReportReq struct { //вспомогательный тип для создания без ID
	ID         int
	DoctorName string
	Diagnosis  string
	IDClient   int
}

func (uc *UseCase) Modify(mod ModifyReportReq) error {
	report := domain.NewMedicalReport(mod.DoctorName, mod.Diagnosis)
	report.SetID(mod.ID)

	report.IDClient = mod.IDClient

	err := uc.medRepo.Modify(report)
	if err != nil {
		return fmt.Errorf("modify report: %v", err)
	}

	return nil
}

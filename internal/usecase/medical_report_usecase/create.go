package medical_report_usecase

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
	client, err := uc.clientRepo.FindByID(req.IDClient)
	if err != nil {
		return fmt.Errorf("clientRepo.FindByID: %v", err)
	}

	report := domain.NewMedicalReport(req.DoctorName, req.Diagnosis, client.ID)

	if err = uc.medRepo.Create(report); err != nil {
		return fmt.Errorf("medRepo.Create: %v", err)
	}

	return nil
}

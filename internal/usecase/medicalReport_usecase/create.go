package medicalReport_usecase

import (
	"fmt"
	"project2/internal/domain"
	"time"
)

type CreateMedicalReportReq struct { //вспомогательный тип для создания без ID
	DoctorName string
	Diagnosis  string
}

func (uc *UseCase) Create(req CreateMedicalReportReq, clientID int) error {
	report := domain.NewMedicalReport(req.DoctorName, req.Diagnosis)

	// Получаем текущее время
	now := time.Now()
	// Форматируем время в нужный формат
	formattedTime := now.Format("02.01.2006 15:04")
	report.CreatedAt = formattedTime

	report.IDClient = clientID

	err := uc.medRepo.Create(report)
	if err != nil {
		return fmt.Errorf("create report: %v", err)
	}
	return nil
}

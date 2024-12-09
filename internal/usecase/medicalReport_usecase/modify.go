package medicalReport_usecase

import (
	"fmt"
	"project2/internal/domain"
	"time"
)

type ModifyReport struct { //вспомогательный тип для создания без ID
	DoctorName string
	Diagnosis  string
	IDClient   int
}

func (uc *UseCase) Modify(id int, mod ModifyReport) error {
	report := domain.NewMedicalReport(mod.DoctorName, mod.Diagnosis)
	report.SetID(id)

	// Получаем текущее время
	now := time.Now()
	// Форматируем время в нужный формат
	formattedTime := now.Format("02.01.2006 15:04")
	report.CreatedAt = formattedTime

	report.IDClient = mod.IDClient

	err := uc.medRepo.Modify(report)
	if err != nil {
		return fmt.Errorf("modify report: %v", err)
	}
	return nil
}

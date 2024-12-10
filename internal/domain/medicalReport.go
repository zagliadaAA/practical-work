package domain

import "time"

type MedicalReport struct {
	ID         int
	DoctorName string
	Diagnosis  string
	CreatedAt  string
	IDClient   int
}

func NewMedicalReport(doctorName, diagnosis string) *MedicalReport {
	// Получаем текущее время
	now := time.Now()
	// Форматируем время в нужный формат
	formattedTime := now.Format("02.01.2006 15:04")

	return &MedicalReport{
		DoctorName: doctorName,
		Diagnosis:  diagnosis,
		CreatedAt:  formattedTime,
	}
}

func (c *MedicalReport) SetID(id int) {
	c.ID = id
}

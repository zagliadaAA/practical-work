package domain

import (
	"fmt"
	"time"
)

type MedicalReport struct {
	ID         int
	DoctorName string
	Diagnosis  string
	CreatedAt  time.Time
	IDClient   int
}

func NewMedicalReport(doctorName, diagnosis string, clientID int) *MedicalReport {
	now := time.Now().UTC()
	dateStr := now.Format("02.01.2006 15:04")
	createdAt, err := time.Parse("02.01.2006 15:04", dateStr)
	if err != nil {
		fmt.Println("не получилось преобразовать дату")
		return nil
	}

	return &MedicalReport{
		DoctorName: doctorName,
		Diagnosis:  diagnosis,
		CreatedAt:  createdAt,
		IDClient:   clientID,
	}
}

func (c *MedicalReport) SetID(id int) {
	c.ID = id
}

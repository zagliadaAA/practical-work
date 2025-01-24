package domain

import (
	"time"
)

type MedicalReport struct {
	ID         int
	DoctorName string
	Diagnosis  string
	CreatedAt  time.Time
	IDClient   int
}

func NewMedicalReport(doctorName, diagnosis string, clientID int, createdAt time.Time) *MedicalReport {
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

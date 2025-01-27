package domain

import (
	"time"
)

type MedicalReport struct {
	ID         int
	DoctorName string
	Diagnosis  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IDClient   int
}

func NewMedicalReport(doctorName, diagnosis string, clientID int, createdAt, updatedAt time.Time) *MedicalReport {
	return &MedicalReport{
		DoctorName: doctorName,
		Diagnosis:  diagnosis,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		IDClient:   clientID,
	}
}

func (c *MedicalReport) SetID(id int) {
	c.ID = id
}

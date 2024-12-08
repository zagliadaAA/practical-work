package domain

type MedicalReport struct {
	ID         int
	DoctorName string
	Diagnosis  string
	CreatedAt  string
	IDClient   int
}

func NewMedicalReport(doctorName, diagnosis string) *MedicalReport {
	return &MedicalReport{
		DoctorName: doctorName,
		Diagnosis:  diagnosis,
	}
}

func (c *MedicalReport) SetID(id int) {
	c.ID = id
}

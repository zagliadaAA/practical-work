package medical_report_repository

import (
	"medicalCenter/internal/domain"
)

func (r *MedRepo) Create(medicalReport *domain.MedicalReport) (*domain.MedicalReport, error) {
	id := r.getNextIdentifier()
	medicalReport.SetID(id)

	r.reportMap[medicalReport.ID] = medicalReport

	return medicalReport, nil
}

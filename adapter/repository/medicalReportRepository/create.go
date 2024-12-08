package medicalReportRepository

import (
	"project2/internal/domain"
)

func (r *MedRepo) Create(medicalReport *domain.MedicalReport) error {

	id := r.getNextIdentifier()
	medicalReport.SetID(id)

	r.reportMap[medicalReport.ID] = medicalReport

	return nil
}

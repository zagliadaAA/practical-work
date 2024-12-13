package medical_report_repository

import (
	"project2/internal/domain"
)

func (r *MedRepo) Update(medicalReport *domain.MedicalReport) error {
	if _, ok := r.reportMap[medicalReport.ID]; !ok {
		return ErrReportNotFound
	}

	r.reportMap[medicalReport.ID] = medicalReport

	return nil
}

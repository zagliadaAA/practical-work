package medical_report_repository

import (
	"project2/internal/domain"
)

func (r *MedRepo) Update(medicalReport *domain.MedicalReport) (*domain.MedicalReport, error) {
	if _, ok := r.reportMap[medicalReport.ID]; !ok {
		return nil, ErrReportNotFound
	}

	r.reportMap[medicalReport.ID] = medicalReport

	return medicalReport, nil
}

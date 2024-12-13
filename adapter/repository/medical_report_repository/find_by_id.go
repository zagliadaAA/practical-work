package medical_report_repository

import (
	"project2/internal/domain"
)

func (r *MedRepo) FindByID(id int) (*domain.MedicalReport, error) {
	report, ok := r.reportMap[id]
	if !ok {
		return nil, ErrReportNotFound
	}

	return report, nil
}

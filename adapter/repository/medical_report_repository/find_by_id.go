package medical_report_repository

import (
	"errors"
	"project2/internal/domain"
)

func (r *MedRepo) FindByID(id int) (*domain.MedicalReport, error) {
	report, ok := r.reportMap[id]
	if !ok {
		return nil, errors.New("report not found")
	}

	return report, nil
}

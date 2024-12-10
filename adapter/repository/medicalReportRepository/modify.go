package medicalReportRepository

import (
	"fmt"
	"project2/internal/domain"
)

func (r *MedRepo) Modify(medicalReport *domain.MedicalReport) error {
	if _, ok := r.reportMap[medicalReport.ID]; !ok {
		return fmt.Errorf("диагноза не существует")
	}

	r.reportMap[medicalReport.ID] = medicalReport

	return nil
}

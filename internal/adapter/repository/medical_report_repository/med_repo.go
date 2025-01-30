package medical_report_repository

import (
	"errors"

	"project2/internal/domain"
)

var ErrReportNotFound = errors.New("report not found")

type MedRepo struct {
	reportMap map[int]*domain.MedicalReport
	id        int
}

func NewMedRepo() *MedRepo {
	return &MedRepo{
		reportMap: make(map[int]*domain.MedicalReport, 100),
		id:        1,
	}
}

func (r *MedRepo) getNextIdentifier() int {
	currentID := r.id
	r.id++

	return currentID
}

func (r *MedRepo) GetAll() []domain.MedicalReport {
	reports := make([]domain.MedicalReport, 0, len(r.reportMap))
	for _, report := range r.reportMap {
		reports = append(reports, *report)
	}
	return reports
}

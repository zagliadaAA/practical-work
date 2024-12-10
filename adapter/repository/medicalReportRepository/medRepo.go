package medicalReportRepository

import (
	"fmt"
	"project2/internal/domain"
)

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

// GetIdReport ищет ID в базе report по ID клиента
func (r *MedRepo) GetIdReport(id int) (int, error) {
	for key, client := range r.reportMap {
		if client.IDClient == id {
			return key, nil
		}
	}

	return 0, fmt.Errorf("клиент с id %d не найден", id)
}

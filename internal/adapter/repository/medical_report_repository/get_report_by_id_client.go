package medical_report_repository

import "medicalCenter/internal/domain"

// GetReportByIDClient ищет и возвращает заключение по id клиента
func (r *MedRepo) GetReportByIDClient(clientID int) (*domain.MedicalReport, error) {
	for i, rep := range r.reportMap {
		if rep.IDClient == clientID {
			report := r.reportMap[i]
			return report, nil
		}
	}

	return nil, ErrReportNotFound
}

package reports

import (
	"context"
	"fmt"

	"project2/internal/domain"
)

func (r *MedRepo) GetReportByIDClient(clientID int) (*domain.MedicalReport, error) {
	query := "SELECT id, doctor_name, diagnosis, created_at, updated_at, id_client FROM reports WHERE id_client = $1;"
	var report domain.MedicalReport

	err := r.cluster.Conn.QueryRow(context.Background(), query, clientID).
		Scan(&report.ID, &report.DoctorName, &report.Diagnosis, &report.CreatedAt, &report.UpdatedAt, &report.IDClient)

	if err != nil {
		return nil, fmt.Errorf("GetReportByIDClient: error finding report: %w", err)
	}

	return &report, nil
}

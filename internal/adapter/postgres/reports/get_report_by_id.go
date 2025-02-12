package reports

import (
	"context"
	"errors"
	"fmt"

	"medicalCenter/internal/adapter/postgres"
	"medicalCenter/internal/domain"

	"github.com/jackc/pgx/v5"
)

func (r *MedRepo) GetReportByID(clientID int) (*domain.MedicalReport, error) {
	query := "SELECT id, doctor_name, diagnosis, created_at, updated_at, id_client FROM reports WHERE id = $1;"
	var report domain.MedicalReport

	err := r.cluster.Conn.QueryRow(context.Background(), query, clientID).
		Scan(&report.ID, &report.DoctorName, &report.Diagnosis, &report.CreatedAt, &report.UpdatedAt, &report.IDClient)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, postgres.ErrNotFound
		}
		return nil, fmt.Errorf("GetReportByID: error finding report: %w", err)
	}

	return &report, nil
}

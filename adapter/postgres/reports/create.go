package reports

import (
	"context"
	"fmt"

	"project2/internal/domain"
)

func (r *MedRepo) Create(mr *domain.MedicalReport) (*domain.MedicalReport, error) {
	query := "INSERT INTO reports(doctor_name, diagnosis, created_at, updated_at, id_client) VALUES ($1, $2, $3, $4, $5) RETURNING id;"

	err := r.cluster.Conn.QueryRow(context.Background(), query, mr.DoctorName, mr.Diagnosis, mr.CreatedAt, mr.UpdatedAt, mr.IDClient).Scan(&mr.ID)

	if err != nil {
		return nil, fmt.Errorf("createReport: error creating report: %w", err)
	}

	return mr, nil
}

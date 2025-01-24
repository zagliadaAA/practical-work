package reports

import (
	"context"
	"fmt"

	"project2/internal/domain"
)

func (r *MedRepo) Update(mr *domain.MedicalReport) (*domain.MedicalReport, error) {
	query := "UPDATE reports SET doctor_name = $1, diagnosis = $2 WHERE id = $3 " +
		"RETURNING id, doctor_name, diagnosis, created_at, id_client;"

	err := r.cluster.Conn.QueryRow(context.Background(), query,
		mr.DoctorName, mr.Diagnosis, mr.ID).Scan(&mr.ID, &mr.DoctorName, &mr.Diagnosis, &mr.CreatedAt, &mr.IDClient)

	if err != nil {
		return nil, fmt.Errorf("updateReport: error updating report: %w", err)
	}

	return mr, nil
}

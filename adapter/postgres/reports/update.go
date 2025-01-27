package reports

import (
	"context"
	"fmt"

	"project2/internal/domain"
)

func (r *MedRepo) Update(mr *domain.MedicalReport) (*domain.MedicalReport, error) {
	query := "UPDATE reports SET doctor_name = $1, diagnosis = $2, updated_at = $3 WHERE id = $4;"

	_, err := r.cluster.Conn.Exec(context.Background(), query, mr.DoctorName, mr.Diagnosis, mr.UpdatedAt, mr.ID)

	if err != nil {
		return nil, fmt.Errorf("updateReport: error updating report: %w", err)
	}

	return mr, nil
}

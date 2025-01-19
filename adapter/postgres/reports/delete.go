package reports

import (
	"context"
	"fmt"
)

func (r *MedRepo) Delete(id int) error {
	query := "DELETE FROM reports WHERE id = $1;"

	mr, err := r.cluster.Conn.Exec(context.Background(), query, id)

	if err != nil {
		return fmt.Errorf("deleteReport: error deleting report: %w", err)
	}

	if mr.RowsAffected() == 0 {
		return fmt.Errorf("report with id %d not found", id)
	}

	return nil
}

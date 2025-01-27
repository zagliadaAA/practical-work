package clients

import (
	"context"
	"fmt"
)

func (r *Repo) Delete(id int) error {
	query := "DELETE FROM clients WHERE id = $1;"

	c, err := r.cluster.Conn.Exec(context.Background(), query, id)

	if err != nil {
		return fmt.Errorf("deleteClient: error deleting client: %w", err)
	}

	if c.RowsAffected() == 0 {
		return fmt.Errorf("client with id %d not found", id)
	}

	return nil
}

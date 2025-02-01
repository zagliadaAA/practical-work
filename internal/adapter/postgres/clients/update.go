package clients

import (
	"context"
	"fmt"

	"medicalCenter/internal/domain"
)

func (r *Repo) Update(c *domain.Client) (*domain.Client, error) {
	query := "UPDATE clients SET name = $1, b_date = $2, phone_number = $3, updated_at = $4 WHERE id = $5;"

	_, err := r.cluster.Conn.Exec(context.Background(), query, c.Name, c.BDate, c.PhoneNumber, c.UpdatedAt, c.ID)

	if err != nil {
		return nil, fmt.Errorf("updateClient: error updating client: %w", err)
	}

	return c, nil
}

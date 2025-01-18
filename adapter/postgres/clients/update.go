package clients

import (
	"context"
	"fmt"

	"project2/internal/domain"
)

func (r *Repo) Update(c *domain.Client) (*domain.Client, error) {
	query := "UPDATE clients SET name = $1, b_date = $2, phone_number = $3 WHERE id = $4 " +
		"RETURNING id, name, b_date, phone_number;"

	err := r.cluster.Conn.QueryRow(context.Background(), query,
		c.Name, c.BDate, c.PhoneNumber, c.ID).Scan(&c.ID, &c.Name, &c.BDate, &c.PhoneNumber)

	if err != nil {
		return nil, fmt.Errorf("updateClient: error updating client: %w", err)
	}

	return c, nil
}

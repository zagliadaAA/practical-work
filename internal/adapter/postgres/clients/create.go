package clients

import (
	"context"
	"fmt"

	"medicalCenter/internal/domain"
)

func (r *Repo) Create(c *domain.Client) (*domain.Client, error) {
	query := "INSERT INTO clients(name, b_date, phone_number, updated_at) VALUES ($1, $2, $3, $4) RETURNING id;"

	err := r.cluster.Conn.QueryRow(context.Background(), query, c.Name, c.BDate, c.PhoneNumber, c.UpdatedAt).Scan(&c.ID)

	if err != nil {
		return nil, fmt.Errorf("createClient: error creating client: %w", err)
	}

	return c, nil
}

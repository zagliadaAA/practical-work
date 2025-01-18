package clients

import (
	"context"
	"fmt"

	"project2/internal/domain"
)

func (r *Repo) Create(c *domain.Client) (*domain.Client, error) {
	query := "INSERT INTO clients(name, b_date, phone_number) VALUES ($1, $2, $3) RETURNING id;"

	err := r.cluster.Conn.QueryRow(context.Background(), query, c.Name, c.BDate, c.PhoneNumber).Scan(&c.ID)

	if err != nil {
		return nil, fmt.Errorf("createClient: error creating client: %w", err)
	}

	return c, nil
}

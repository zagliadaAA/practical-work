package clients

import (
	"context"
	"fmt"

	"project2/internal/domain"
)

func (r *Repo) FindByID(id int) (*domain.Client, error) {
	query := "SELECT id, name, b_date, phone_number, updated_at FROM clients WHERE id = $1;"

	var client domain.Client

	err := r.cluster.Conn.QueryRow(context.Background(), query, id).Scan(&client.ID, &client.Name, &client.BDate, &client.PhoneNumber, &client.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("findByID: error finding client: %w", err)
	}

	return &client, nil
}

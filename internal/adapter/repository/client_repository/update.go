package client_repository

import (
	"medicalCenter/internal/domain"
)

func (r *Repo) Update(client *domain.Client) (*domain.Client, error) {
	if _, ok := r.clientMap[client.ID]; !ok {
		return nil, ErrClientNotFound
	}

	r.clientMap[client.ID] = client

	return client, nil
}

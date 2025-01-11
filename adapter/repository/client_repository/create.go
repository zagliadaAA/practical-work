package client_repository

import "project2/internal/domain"

func (r *Repo) Create(client *domain.Client) (*domain.Client, error) {
	id := r.getNextIdentifier()
	client.SetID(id)

	r.clientMap[client.ID] = client

	return client, nil
}

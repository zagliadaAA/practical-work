package client_repository

import (
	"errors"
	"project2/internal/domain"
)

func (r *Repo) FindByID(id int) (*domain.Client, error) {
	client, ok := r.clientMap[id]
	if !ok {
		return nil, errors.New("client not found")
	}

	return client, nil
}

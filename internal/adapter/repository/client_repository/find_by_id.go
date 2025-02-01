package client_repository

import (
	"medicalCenter/internal/domain"
)

func (r *Repo) FindByID(id int) (*domain.Client, error) {
	client, ok := r.clientMap[id]
	if !ok {
		return nil, ErrClientNotFound
	}

	return client, nil
}

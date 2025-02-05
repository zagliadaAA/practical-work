package client_usecase

import (
	"fmt"

	"medicalCenter/internal/domain"
)

func (uc *UseCase) GetClientByID(id int) (*domain.Client, error) {
	client, err := uc.clientRepo.GetClientByID(id)
	if err != nil {
		return nil, fmt.Errorf("clientRepo.GetClientByID: %w", err)
	}

	return client, nil
}

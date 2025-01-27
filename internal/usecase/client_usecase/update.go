package client_usecase

import (
	"fmt"
	"time"

	"project2/internal/domain"
)

type UpdateClientReq struct { //Вспомогательный тип для создания без ID
	ID          int
	Name        string
	BDate       time.Time
	PhoneNumber string
}

func (uc *UseCase) Update(req UpdateClientReq) (*domain.Client, error) {
	client, err := uc.clientRepo.FindByID(req.ID)
	if err != nil {
		return nil, fmt.Errorf("clientRepo.FindByID: %w", err)
	}

	client.Name = req.Name
	client.BDate = req.BDate
	client.PhoneNumber = req.PhoneNumber
	client.UpdatedAt = uc.timer.Now()

	clientUpdate, err := uc.clientRepo.Update(client)
	if err != nil {
		return nil, fmt.Errorf("clientRepo.Update: %w", err)
	}

	return clientUpdate, nil
}

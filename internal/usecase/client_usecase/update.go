package client_usecase

import (
	"fmt"
	"time"

	"project2/internal/domain"
)

type UpdateClientReq struct { //Вспомогательный тип для создания без ID
	ID          int
	Name        string
	BDate       string
	PhoneNumber string
}

func (uc *UseCase) Update(req UpdateClientReq) (*domain.Client, error) {
	BDate, err := time.Parse("02.01.2006", req.BDate)
	if err != nil {
		return nil, fmt.Errorf("date conversion error: %w", err)
	}

	client, err := uc.clientRepo.FindByID(req.ID)
	if err != nil {
		return nil, fmt.Errorf("clientRepo.FindByID: %w", err)
	}

	client.Name = req.Name
	client.BDate = BDate
	client.PhoneNumber = req.PhoneNumber

	clientUpdate, err := uc.clientRepo.Update(client)
	if err != nil {
		return nil, fmt.Errorf("clientRepo.Update: %w", err)
	}

	return clientUpdate, nil
}

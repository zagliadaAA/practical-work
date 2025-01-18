package client_usecase

import (
	"fmt"
	"time"

	"project2/internal/domain"
)

type CreateClientReq struct { //Вспомогательный тип для создания без ID
	Name        string
	BDate       string
	PhoneNumber string
}

func (uc *UseCase) Create(req CreateClientReq) (*domain.Client, error) {
	BDate, err := time.Parse("02.01.2006", req.BDate)
	if err != nil {
		return nil, fmt.Errorf("date conversion error: %w", err)
	}

	client := domain.NewClient(req.Name, BDate, req.PhoneNumber)

	clientCreate, err := uc.clientRepo.Create(client)
	if err != nil {
		return nil, fmt.Errorf("clientRepo.Create: %w", err)
	}

	return clientCreate, nil
}

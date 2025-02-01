package client_usecase

import (
	"fmt"
	"time"

	"medicalCenter/internal/domain"
)

type CreateClientReq struct { //Вспомогательный тип для создания без ID
	Name        string
	BDate       time.Time
	PhoneNumber string
}

func (uc *UseCase) Create(req CreateClientReq) (*domain.Client, error) {
	client := domain.NewClient(req.Name, req.BDate, req.PhoneNumber, uc.timer.Now())

	clientCreate, err := uc.clientRepo.Create(client)
	if err != nil {
		return nil, fmt.Errorf("clientRepo.Create: %w", err)
	}

	return clientCreate, nil
}

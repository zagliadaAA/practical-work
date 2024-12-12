package client_usecase

import (
	"fmt"
	"project2/internal/domain"
)

type CreateClientReq struct { //вспомогательный тип для создания без ID
	Name        string
	BDate       string
	PhoneNumber string
}

func (uc *UseCase) Create(req CreateClientReq) error {
	client := domain.NewClient(req.Name, req.BDate, req.PhoneNumber)

	if err := uc.clientRepo.Create(client); err != nil {
		return fmt.Errorf("create client: %v", err)
	}

	return nil
}

package client_usecase

import (
	"fmt"
)

type UpdateClientReq struct { //вспомогательный тип для создания без ID
	ID          int
	Name        string
	BDate       string
	PhoneNumber string
}

func (uc *UseCase) Update(req UpdateClientReq) error {
	client, err := uc.clientRepo.FindByID(req.ID)
	if err != nil {
		return fmt.Errorf("clientRepo.FindByID: %w", err)
	}

	client.Name = req.Name
	client.BDate = req.BDate
	client.PhoneNumber = req.PhoneNumber

	if err = uc.clientRepo.Update(client); err != nil {
		return fmt.Errorf("clientRepo.Update: %w", err)
	}

	return nil
}

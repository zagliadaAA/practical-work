package client_usecase

import (
	"fmt"
	"project2/internal/domain"
)

type ModifyClientReq struct { //вспомогательный тип для создания без ID
	ID          int
	Name        string
	BDate       string
	PhoneNumber string
}

func (uc *UseCase) Modify(mod ModifyClientReq) error {
	client := domain.NewClient(mod.Name, mod.BDate, mod.PhoneNumber)
	client.SetID(mod.ID)
	err := uc.clientRepo.Modify(client)
	if err != nil {
		return fmt.Errorf("modify client: %v", err)
	}

	return nil
}

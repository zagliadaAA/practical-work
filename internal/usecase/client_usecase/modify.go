package client_usecase

import (
	"fmt"
	"project2/internal/domain"
)

type ModifyClient struct { //вспомогательный тип для создания без ID
	Name        string
	BDate       string
	PhoneNumber string
}

func (uc *UseCase) Modify(id int, mod CreateClientReq) error {
	client := domain.NewClient(mod.Name, mod.BDate, mod.PhoneNumber)
	client.SetID(id)
	err := uc.clientRepo.Modify(client)
	if err != nil {
		return fmt.Errorf("modify client: %v", err)
	}
	return nil
}

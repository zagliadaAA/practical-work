package client_usecase

import (
	"fmt"
)

func (uc *UseCase) Delete(id int) error {
	err := uc.clientRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("delete client: %v", err)
	}
	return nil
}

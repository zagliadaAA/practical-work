package client_usecase

import (
	"fmt"
)

func (uc *UseCase) Delete(id int) error {
	if err := uc.clientRepo.Delete(id); err != nil {
		return fmt.Errorf("delete client: %v", err)
	}

	return nil
}

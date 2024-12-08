package clientRepository

import (
	"fmt"
	"project2/internal/domain"
)

func (r *Repo) Modify(client *domain.Client) error {

	if _, ok := r.clientMap[client.ID]; !ok {
		return fmt.Errorf("клиента не существует")
	}

	r.clientMap[client.ID] = client

	return nil
}

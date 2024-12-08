package clientRepository

import "fmt"

func (r *Repo) Delete(id int) error {

	if _, ok := r.clientMap[id]; !ok {
		return fmt.Errorf("клиента не существует")
	}

	delete(r.clientMap, id)

	return nil
}

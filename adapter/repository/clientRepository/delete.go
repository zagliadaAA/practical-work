package clientRepository

func (r *Repo) Delete(id int) error {
	if _, ok := r.clientMap[id]; !ok {
		return nil
	}

	delete(r.clientMap, id)

	return nil
}

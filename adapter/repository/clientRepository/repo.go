package clientRepository

import (
	"project2/internal/domain"
)

type Repo struct {
	clientMap map[int]*domain.Client
	id        int
}

func NewRepo() *Repo {
	return &Repo{
		clientMap: make(map[int]*domain.Client, 100),
		id:        1,
	}
}

func (r *Repo) getNextIdentifier() int {
	currentID := r.id
	r.id++

	return currentID
}

func (r *Repo) GetAll() []domain.Client {
	clients := make([]domain.Client, 0, len(r.clientMap))
	for _, client := range r.clientMap {
		clients = append(clients, *client)
	}
	return clients
}

// FindID существует ли клиент в базе
func (r *Repo) FindID(id int) bool {
	if _, ok := r.clientMap[id]; !ok {
		return false
	}

	return true
}

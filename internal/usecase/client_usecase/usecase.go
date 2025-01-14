package client_usecase

import "project2/internal/domain"

type UseCase struct {
	clientRepo clientRepo
}

type clientRepo interface {
	Create(client *domain.Client) (*domain.Client, error)
	Delete(id int) error
	Update(client *domain.Client) (*domain.Client, error)
	FindByID(id int) (*domain.Client, error)
}

func NewUseCase(clientRepo clientRepo) *UseCase {
	return &UseCase{
		clientRepo: clientRepo,
	}
}

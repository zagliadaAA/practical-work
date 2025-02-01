package client_usecase

import (
	"time"

	"medicalCenter/internal/domain"
)

type UseCase struct {
	clientRepo clientRepo
	timer      timer
}

type (
	clientRepo interface {
		Create(client *domain.Client) (*domain.Client, error)
		Delete(id int) error
		Update(client *domain.Client) (*domain.Client, error)
		FindByID(id int) (*domain.Client, error)
	}
	timer interface {
		Now() time.Time
	}
)

func NewUseCase(clientRepo clientRepo, timer timer) *UseCase {
	return &UseCase{
		clientRepo: clientRepo,
		timer:      timer,
	}
}

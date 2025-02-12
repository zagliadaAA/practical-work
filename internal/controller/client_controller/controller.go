package client_controller

import (
	"medicalCenter/internal/domain"
	"medicalCenter/internal/usecase/client_usecase"
)

type Controller struct {
	clientUseCase clientUseCase
}

type clientUseCase interface {
	Create(req client_usecase.CreateClientReq) (*domain.Client, error)
	Delete(id int) error
	Update(req client_usecase.UpdateClientReq) (*domain.Client, error)
	GetClientByID(id int) (*domain.Client, error)
}

func NewController(useCase clientUseCase) *Controller {
	return &Controller{
		clientUseCase: useCase,
	}
}

package client_controller

import (
	"medicalCenter/internal/domain"
	"medicalCenter/internal/usecase/client_usecase"
)

type (
	Controller struct {
		clientUseCase clientUseCase
	}

	clientUseCase interface {
		Create(req client_usecase.CreateClientReq) (*domain.Client, error)
	}
)

func NewController(useCase clientUseCase) *Controller {
	return &Controller{
		clientUseCase: useCase,
	}
}

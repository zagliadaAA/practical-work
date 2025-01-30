package client_controller

import (
	"project2/internal/domain"
	"project2/internal/usecase/client_usecase"
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

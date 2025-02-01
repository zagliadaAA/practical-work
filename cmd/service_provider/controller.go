package service_provider

import (
	"medicalCenter/internal/controller/client_controller"
)

func (sp *ServiceProvider) getClientController() *client_controller.Controller {
	if sp.clientController == nil {
		sp.clientController = client_controller.NewController(sp.GetClientUseCase())
	}

	return sp.clientController
}

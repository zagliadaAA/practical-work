package service_provider

import (
	"medicalCenter/internal/controller/client_controller"
	"medicalCenter/internal/controller/medical_report_controller"
)

func (sp *ServiceProvider) getClientController() *client_controller.Controller {
	if sp.clientController == nil {
		sp.clientController = client_controller.NewController(sp.GetClientUseCase())
	}

	return sp.clientController
}

func (sp *ServiceProvider) getMedicalReportController() *medical_report_controller.Controller {
	if sp.medicalReportController == nil {
		sp.medicalReportController = medical_report_controller.NewController(sp.GetMedicalReportUseCase())
	}

	return sp.medicalReportController
}

package service_provider

import (
	"medicalCenter/internal/usecase/client_usecase"
	"medicalCenter/internal/usecase/medical_report_usecase"
)

func (sp *ServiceProvider) GetClientUseCase() *client_usecase.UseCase {
	if sp.clientUseCase == nil {
		sp.clientUseCase = client_usecase.NewUseCase(sp.GetClientRepository(), sp.getTimer())
	}

	return sp.clientUseCase
}

func (sp *ServiceProvider) GetMedicalReportUseCase() *medical_report_usecase.UseCase {
	if sp.medicalReportUseCase == nil {
		sp.medicalReportUseCase = medical_report_usecase.NewUseCase(sp.GetMedicalReportRepository(), sp.GetClientRepository(), sp.getTimer())
	}

	return sp.medicalReportUseCase
}

package service_provider

import (
	"project2/adapter/repository/client_repository"
	"project2/adapter/repository/medical_report_repository"
)

func (sp *ServiceProvider) GetClientRepository() *client_repository.Repo {
	if sp.clientRepo == nil {
		sp.clientRepo = client_repository.NewRepo()
	}

	return sp.clientRepo
}

func (sp *ServiceProvider) GetMedicalReportRepository() *medical_report_repository.MedRepo {
	if sp.medicalReportRepo == nil {
		sp.medicalReportRepo = medical_report_repository.NewMedRepo()
	}

	return sp.medicalReportRepo
}

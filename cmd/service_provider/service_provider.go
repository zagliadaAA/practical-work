package service_provider

import (
	"project2/adapter/postgres/clients"
	"project2/adapter/repository/medical_report_repository"
	"project2/internal/config"
	"project2/internal/usecase/client_usecase"
	"project2/internal/usecase/medical_report_usecase"
)

type ServiceProvider struct {
	dbCluster *config.Cluster

	clientUseCase        *client_usecase.UseCase
	medicalReportUseCase *medical_report_usecase.UseCase

	//clientRepo        *client_repository.Repo
	clientRepo        *clients.Repo
	medicalReportRepo *medical_report_repository.MedRepo
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

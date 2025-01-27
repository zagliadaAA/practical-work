package service_provider

import (
	"project2/adapter/postgres/clients"
	"project2/adapter/postgres/reports"
	"project2/internal/config"
	"project2/internal/pkg/timer"
	"project2/internal/usecase/client_usecase"
	"project2/internal/usecase/medical_report_usecase"
)

type ServiceProvider struct {
	dbCluster *config.Cluster

	clientUseCase        *client_usecase.UseCase
	medicalReportUseCase *medical_report_usecase.UseCase

	clientRepo        *clients.Repo
	medicalReportRepo *reports.MedRepo

	timer *timer.Timer
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

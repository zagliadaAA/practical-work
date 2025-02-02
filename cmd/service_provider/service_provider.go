package service_provider

import (
	"medicalCenter/internal/adapter/postgres/clients"
	"medicalCenter/internal/adapter/postgres/reports"
	"medicalCenter/internal/config"
	"medicalCenter/internal/controller/client_controller"
	"medicalCenter/internal/controller/medical_report_controller"
	"medicalCenter/internal/pkg/timer"
	"medicalCenter/internal/usecase/client_usecase"
	"medicalCenter/internal/usecase/medical_report_usecase"
)

type ServiceProvider struct {
	dbCluster *config.Cluster

	clientUseCase        *client_usecase.UseCase
	medicalReportUseCase *medical_report_usecase.UseCase

	clientRepo        *clients.Repo
	medicalReportRepo *reports.MedRepo

	clientController        *client_controller.Controller
	medicalReportController *medical_report_controller.Controller

	timer *timer.Timer
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

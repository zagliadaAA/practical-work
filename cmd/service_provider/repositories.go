package service_provider

import (
	"context"

	"medicalCenter/internal/adapter/postgres/clients"
	"medicalCenter/internal/adapter/postgres/reports"
)

func (sp *ServiceProvider) GetClientRepository() *clients.Repo {
	if sp.clientRepo == nil {
		sp.clientRepo = clients.NewRepo(sp.GetDbCluster(context.Background()))
	}

	return sp.clientRepo
}

func (sp *ServiceProvider) GetMedicalReportRepository() *reports.MedRepo {
	if sp.medicalReportRepo == nil {
		sp.medicalReportRepo = reports.NewMedRepo(sp.GetDbCluster(context.Background()))
	}

	return sp.medicalReportRepo
}

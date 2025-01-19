package service_provider

import (
	"context"

	"project2/adapter/postgres/clients"
	"project2/adapter/postgres/reports"
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

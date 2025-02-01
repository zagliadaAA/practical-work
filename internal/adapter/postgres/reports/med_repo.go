package reports

import (
	"medicalCenter/internal/config"
)

type MedRepo struct {
	cluster *config.Cluster
}

func NewMedRepo(cluster *config.Cluster) *MedRepo {
	return &MedRepo{
		cluster: cluster,
	}
}

package reports

import (
	"project2/internal/config"
)

type MedRepo struct {
	cluster *config.Cluster
}

func NewMedRepo(cluster *config.Cluster) *MedRepo {
	return &MedRepo{
		cluster: cluster,
	}
}

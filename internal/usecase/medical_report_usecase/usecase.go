package medical_report_usecase

import "project2/internal/domain"

type UseCase struct {
	medRepo    medRepo
	clientRepo clientRepo
}

type medRepo interface {
	Create(medicalReport *domain.MedicalReport) error
	Delete(id int) error
	Update(medicalReport *domain.MedicalReport) (*domain.MedicalReport, error)
	GetReportByIDClient(clientID int) (*domain.MedicalReport, error)
}

type clientRepo interface {
	FindByID(id int) (*domain.Client, error)
}

func NewUseCase(medRepo medRepo, clientRepo clientRepo) *UseCase {
	return &UseCase{
		medRepo:    medRepo,
		clientRepo: clientRepo,
	}
}

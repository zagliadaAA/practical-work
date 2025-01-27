package medical_report_usecase

import (
	"time"

	"project2/internal/domain"
)

type UseCase struct {
	medRepo    medRepo
	clientRepo clientRepo
	timer      timer
}

type (
	medRepo interface {
		Create(medicalReport *domain.MedicalReport) (*domain.MedicalReport, error)
		Delete(id int) error
		Update(medicalReport *domain.MedicalReport) (*domain.MedicalReport, error)
		GetReportByIDClient(clientID int) (*domain.MedicalReport, error)
	}
	timer interface {
		Now() time.Time
	}
	clientRepo interface {
		FindByID(id int) (*domain.Client, error)
	}
)

func NewUseCase(medRepo medRepo, clientRepo clientRepo, timer timer) *UseCase {
	return &UseCase{
		medRepo:    medRepo,
		clientRepo: clientRepo,
		timer:      timer,
	}
}

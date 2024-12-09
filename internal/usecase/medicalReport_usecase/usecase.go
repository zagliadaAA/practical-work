package medicalReport_usecase

import "project2/internal/domain"

type UseCase struct {
	medRepo medRepo
}

type medRepo interface {
	Create(medicalReport *domain.MedicalReport) error
	Delete(id int) error
	Modify(medicalReport *domain.MedicalReport) error
}

func NewUseCase(medRepo medRepo) *UseCase {
	return &UseCase{
		medRepo: medRepo,
	}
}

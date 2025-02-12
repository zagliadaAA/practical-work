package medical_report_usecase

import "fmt"

func (uc *UseCase) Delete(id int) error {
	if err := uc.medRepo.Delete(id); err != nil {
		return fmt.Errorf("medRepo.Delete: %w", err)
	}

	return nil
}

package medical_report_usecase

import "fmt"

func (uc *UseCase) Delete(id int) error {
	reportID, _ := uc.medRepo.GetIdReport(id)

	if err := uc.medRepo.Delete(reportID); err != nil {
		return fmt.Errorf("medRepo.Delete: %w", err)
	}

	return nil
}

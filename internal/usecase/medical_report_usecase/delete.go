package medical_report_usecase

import "fmt"

func (uc *UseCase) Delete(id int) error {
	report, err := uc.medRepo.GetReportByIDClient(id)
	if err != nil {
		return nil
	}

	if err = uc.medRepo.Delete(report.ID); err != nil {
		return fmt.Errorf("medRepo.Delete: %w", err)
	}

	return nil
}

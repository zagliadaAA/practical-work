package medicalReport_usecase

import "fmt"

func (uc *UseCase) Delete(id int) error {
	err := uc.medRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("delete report: %v", err)
	}
	return nil
}

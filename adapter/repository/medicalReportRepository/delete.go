package medicalReportRepository

import "fmt"

func (r *MedRepo) Delete(id int) error {

	if _, ok := r.reportMap[id]; !ok {
		return fmt.Errorf("диагноза не существует")
	}

	delete(r.reportMap, id)

	return nil
}

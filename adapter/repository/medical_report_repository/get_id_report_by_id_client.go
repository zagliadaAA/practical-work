package medical_report_repository

import "fmt"

// GetIdReport ищет ID в базе report по ID клиента
func (r *MedRepo) GetIdReport(id int) (int, error) {
	for key, client := range r.reportMap {
		if client.IDClient == id {
			return key, nil
		}
	}

	return 0, fmt.Errorf("клиент с id %d не найден", id)
}

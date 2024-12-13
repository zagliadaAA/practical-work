package medical_report_repository

// GetIdReport ищет ID в базе report по ID клиента
func (r *MedRepo) GetIdReport(id int) (int, error) {
	for key, client := range r.reportMap {
		if client.IDClient == id {
			return key, nil
		}
	}

	return 0, ErrReportNotFound
}

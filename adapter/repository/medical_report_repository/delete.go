package medical_report_repository

func (r *MedRepo) Delete(id int) error {
	if _, ok := r.reportMap[id]; !ok {
		return nil
	}

	delete(r.reportMap, id)

	return nil
}

package medical_report_controller

import (
	"net/http"
	"strconv"
)

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Path[len("/reports/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)

		return
	}

	err = c.medicalReportUseCase.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Report deleted successfully"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

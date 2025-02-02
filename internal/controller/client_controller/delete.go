package client_controller

import (
	"net/http"
	"strconv"
)

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Path[len("/clients/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)

		return
	}

	err = c.clientUseCase.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Client deleted successfully"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

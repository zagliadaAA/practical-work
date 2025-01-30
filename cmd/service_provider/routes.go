package service_provider

import (
	"net/http"
)

func (sp *ServiceProvider) GetRoutes() *http.ServeMux {
	// роутер наших ручек
	mux := http.NewServeMux()
	// создание клиента
	mux.HandleFunc("POST /clients", sp.getClientController().Create)
	//mux.HandleFunc("DELETE /clients/{id}", sp.getClientController().Delete)

	return mux
}

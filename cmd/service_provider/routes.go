package service_provider

import (
	"net/http"
)

func (sp *ServiceProvider) GetRoutes() *http.ServeMux {
	// роутер наших ручек (пути для ручек(хэндлеров))
	mux := http.NewServeMux()
	// клиент
	mux.HandleFunc("POST /clients", sp.getClientController().Create)
	mux.HandleFunc("DELETE /clients/{id}", sp.getClientController().Delete)
	mux.HandleFunc("PUT /clients/{id}", sp.getClientController().Update)
	// медицинское заключение
	mux.HandleFunc("POST /reports", sp.getMedicalReportController().Create)
	mux.HandleFunc("DELETE /reports/{id}", sp.getMedicalReportController().Delete)
	mux.HandleFunc("PUT /reports", sp.getMedicalReportController().Update)

	return mux
}

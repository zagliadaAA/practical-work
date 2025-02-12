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
	mux.HandleFunc("GET /clients/{id}", sp.getClientController().GetClientByID)
	// медицинское заключение
	mux.HandleFunc("POST /reports", sp.getMedicalReportController().Create)
	mux.HandleFunc("DELETE /reports/{id}", sp.getMedicalReportController().Delete)
	mux.HandleFunc("PUT /reports/{id}", sp.getMedicalReportController().Update)
	mux.HandleFunc("GET /reports/{id}", sp.getMedicalReportController().GetReportByID)

	return mux
}

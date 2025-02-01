package main

import (
	"net/http"

	"medicalCenter/cmd/service_provider"
)

func main() {
	sp := service_provider.NewServiceProvider()

	// запуск сервера
	err := http.ListenAndServe(":8080", sp.GetRoutes())
	if err != nil {
		panic(err)
	}
}

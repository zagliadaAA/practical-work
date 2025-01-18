package main

import (
	"errors"
	"fmt"

	"project2/adapter/repository/client_repository"
	"project2/cmd/service_provider"
	"project2/internal/usecase/client_usecase"
)

func main() {

	fmt.Println("-----\n")

	sp := service_provider.NewServiceProvider()

	//создание клиента
	_, err := sp.GetClientUseCase().Create(client_usecase.CreateClientReq{
		Name:        "Artem",
		BDate:       "30.12.1999",
		PhoneNumber: "89085538251",
	})
	if err != nil {
		panic(err)
	}

	//изменение клиента
	_, err = sp.GetClientUseCase().Update(client_usecase.UpdateClientReq{
		ID:          22,
		Name:        "Vika",
		BDate:       "30.12.1999",
		PhoneNumber: "89085538251",
	})
	if err != nil {
		if errors.Is(err, client_repository.ErrClientNotFound) {
			fmt.Println("client does not exist")
		} else {
			panic(err)
		}
	}

	//удаление клиента
	err = sp.GetClientUseCase().Delete(21)
	if err != nil {
		panic(err)
	}

	//добавление диагноза для клиента
	/*_, err = sp.GetMedicalReportUseCase().Create(medical_report_usecase.CreateMedicalReportReq{
		IDClient:   1,
		DoctorName: "Доктор Вася",
		Diagnosis:  "F20.5",
	})
	if err != nil {
		if errors.Is(err, client_repository.ErrClientNotFound) {
			fmt.Println("create report for client: client does not exist")
		} else {
			panic(err)
		}
	}

	//добавление диагноза для клиента
	_, err = sp.GetMedicalReportUseCase().Create(medical_report_usecase.CreateMedicalReportReq{
		IDClient:   3,
		DoctorName: "Доктор Вася",
		Diagnosis:  "F20.2",
	})
	if err != nil {
		if errors.Is(err, client_repository.ErrClientNotFound) {
			fmt.Println("create report for client: client does not exist")
		} else {
			panic(err)
		}
	}

	// Печатаем все диагнозы
	reports := sp.GetMedicalReportRepository().GetAll()
	fmt.Println(reports)

	//удаление диагноза для клиента
	err = sp.GetMedicalReportUseCase().Delete(1)
	if err != nil {
		panic(err)
	}

	//изменение диагноза для клиента
	_, err = sp.GetMedicalReportUseCase().Update(medical_report_usecase.UpdateReportReq{
		DoctorName: "Доктор Вася",
		Diagnosis:  "F20.7",
		IDClient:   3,
	})
	if err != nil {
		if errors.Is(err, medical_report_repository.ErrReportNotFound) {
			fmt.Println("report does not exist for this client")
		} else {
			panic(err)
		}
	}

	// Печатаем все диагнозы
	reports = sp.GetMedicalReportRepository().GetAll()
	fmt.Println(reports)*/
}

/*
ws := &CarWashStation
honda := &Honda{}
ws.Wash(honda)

type CarWashStation struct {}

func(s *CarWashStation) Wash(car Car) {
	fmt.Println("Моем машину:")
}

type Car interface {
	BeepBeep() string
}

type Honda struct {}
type Chevrolet struct {}

func(c *Honda) BeepBeep() string {
	return fmt.Sprint("honda")
}

func(c *Chevrolet) BeepBeep() string {
	return fmt.Sprint("chevrolet")
}
*/

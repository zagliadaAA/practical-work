package main

import (
	"fmt"
	"project2/adapter/repository/clientRepository"
	"project2/adapter/repository/medicalReportRepository"
	"project2/internal/usecase/client_usecase"
	"project2/internal/usecase/medicalReport_usecase"
)

func main() {

	fmt.Println("-----\n")

	repo := clientRepository.NewRepo() //создаем новое хранилище которое содержит мапу типа клиент + ID
	uc := client_usecase.NewUseCase(repo)
	medicalRepo := medicalReportRepository.NewMedRepo()
	medUc := medicalReport_usecase.NewUseCase(medicalRepo)

	//создание клиента
	err := uc.Create(client_usecase.CreateClientReq{
		Name:        "Artem",
		BDate:       "30.12.1999",
		PhoneNumber: "89085538251",
	})
	if err != nil {
		panic(err)
	}

	//создание клиента
	err = uc.Create(client_usecase.CreateClientReq{
		Name:        "Boba",
		BDate:       "01.01.1999",
		PhoneNumber: "89087732819",
	})
	if err != nil {
		panic(err)
	}

	//создание клиента
	err = uc.Create(client_usecase.CreateClientReq{
		Name:        "Liza",
		BDate:       "02.02.1998",
		PhoneNumber: "8908846195",
	})
	if err != nil {
		panic(err)
	}

	//печать всех клиентов
	clients := repo.GetAll()
	fmt.Println(clients)

	//изменение клиента
	err = uc.Modify(client_usecase.ModifyClientReq{
		ID:          1,
		Name:        "AArtem",
		BDate:       "30.12.1999",
		PhoneNumber: "89085538251",
	})
	if err != nil {
		panic(err)
	}

	//удаление клиента
	err = uc.Delete(2)
	if err != nil {
		panic(err)
	}

	//добавление диагноза для клиента
	clientID := 1
	err = medUc.Create(medicalReport_usecase.CreateMedicalReportReq{
		IDClient:   clientID,
		DoctorName: "Доктор Вася",
		Diagnosis:  "F20.5",
	})
	if err != nil {
		panic(err)
	}

	//добавление диагноза для клиента
	clientID = 3
	err = medUc.Create(medicalReport_usecase.CreateMedicalReportReq{
		IDClient:   clientID,
		DoctorName: "Доктор Вася",
		Diagnosis:  "F20.2",
	})
	if err != nil {
		panic(err)
	}

	// Печатаем все диагнозы
	reports := medicalRepo.GetAll()
	fmt.Println(reports)

	//удаление диагноза для клиента
	reportID, _ := medicalRepo.GetIdReport(3)
	err = medUc.Delete(reportID)
	if err != nil {
		panic(err)
	}

	//изменение диагноза для клиента
	clientID = 1
	reportID, err = medicalRepo.GetIdReport(clientID)
	if err != nil {
		panic(err)
	}
	err = medUc.Modify(medicalReport_usecase.ModifyReportReq{
		ID:         reportID,
		DoctorName: "Доктор Вася",
		Diagnosis:  "F20.7",
		IDClient:   clientID,
	})
	if err != nil {
		panic(err)
	}

	// Печатаем всех клиентов
	clients = repo.GetAll()
	fmt.Println(clients)

	// Печатаем все диагнозы
	reports = medicalRepo.GetAll()
	fmt.Println(reports)

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

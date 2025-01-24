package main

import (
	"errors"
	"fmt"
	"time"

	"project2/adapter/repository/medical_report_repository"
	"project2/cmd/service_provider"
	"project2/internal/usecase/medical_report_usecase"
)

func main() {

	fmt.Println("-----\n")

	sp := service_provider.NewServiceProvider()

	//создание клиента
	/*bDateClient, err := convertDate("11.11.1800")
	if err != nil {
		panic(err)
	}
	_, err = sp.GetClientUseCase().Create(client_usecase.CreateClientReq{
		Name:        "Lyda",
		BDate:       bDateClient,
		PhoneNumber: "89086627184",
	})
	if err != nil {
		panic(err)
	}*/

	//изменение клиента
	/*bDateClient, err := convertDate("12.12.2000")
	if err != nil {
		panic(err)
	}
	_, err = sp.GetClientUseCase().Update(client_usecase.UpdateClientReq{
		ID:          23,
		Name:        "Lena",
		BDate:       bDateClient,
		PhoneNumber: "89086338251",
	})
	if err != nil {
		if errors.Is(err, client_repository.ErrClientNotFound) {
			fmt.Println("client does not exist")
		} else {
			panic(err)
		}
	}*/

	//удаление клиента
	/*err := sp.GetClientUseCase().Delete(23)
	if err != nil {
		panic(err)
	}*/

	//добавление диагноза для клиента
	/*_, err := sp.GetMedicalReportUseCase().Create(medical_report_usecase.CreateMedicalReportReq{
		IDClient:   26,
		DoctorName: "Доктор Вася",
		Diagnosis:  "F100.5",
	})
	if err != nil {
		if errors.Is(err, client_repository.ErrClientNotFound) {
			fmt.Println("create report for client: client does not exist")
		} else {
			panic(err)
		}
	}*/

	//удаление диагноза для клиента
	/*err := sp.GetMedicalReportUseCase().Delete(25)
	if err != nil {
		panic(err)
	}*/

	//изменение диагноза для клиента
	_, err := sp.GetMedicalReportUseCase().Update(medical_report_usecase.UpdateReportReq{
		DoctorName: "Доктор Вася",
		Diagnosis:  "F99",
		IDClient:   26,
	})
	if err != nil {
		if errors.Is(err, medical_report_repository.ErrReportNotFound) {
			fmt.Println("report does not exist for this client")
		} else {
			panic(err)
		}
	}
}

// Функция для конвертации даты
func convertDate(dateStr string) (time.Time, error) {
	parsedDate, err := time.Parse("02.01.2006", dateStr)
	if err != nil {
		return parsedDate, fmt.Errorf("failed to parse date: %w", err)
	}

	return parsedDate, nil
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

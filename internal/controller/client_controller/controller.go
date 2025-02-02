package client_controller

import (
	"encoding/json"
	"net/http"

	"medicalCenter/internal/domain"
	"medicalCenter/internal/usecase/client_usecase"
)

type Controller struct {
	clientUseCase clientUseCase
}

type clientUseCase interface {
	Create(req client_usecase.CreateClientReq) (*domain.Client, error)
	Delete(id int) error
	Update(req client_usecase.UpdateClientReq) (*domain.Client, error)
}

func NewController(useCase clientUseCase) *Controller {
	return &Controller{
		clientUseCase: useCase,
	}
}

func (c *Controller) DecodeRequest(w http.ResponseWriter, r *http.Request, req interface{}) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return err
	}

	return nil
}

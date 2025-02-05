package client_controller

import (
	"medicalCenter/internal/domain"
)

func mapClientToResponseForCreate(client *domain.Client) *createClientResp {
	return &createClientResp{
		ID:          client.ID,
		Name:        client.Name,
		BDate:       client.BDate,
		PhoneNumber: client.PhoneNumber,
		UpdatedAt:   client.UpdatedAt,
	}
}

func mapClientToResponseForUpdate(client *domain.Client) *updateClientResp {
	return &updateClientResp{
		ID:          client.ID,
		Name:        client.Name,
		BDate:       client.BDate,
		PhoneNumber: client.PhoneNumber,
		UpdatedAt:   client.UpdatedAt,
	}
}

func mapClientToResponseForGetClientByID(client *domain.Client) *getClientByIDResp {
	return &getClientByIDResp{
		ID:          client.ID,
		Name:        client.Name,
		BDate:       client.BDate,
		PhoneNumber: client.PhoneNumber,
		UpdatedAt:   client.UpdatedAt,
	}
}

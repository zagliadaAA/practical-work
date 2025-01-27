package domain

import (
	"time"
)

type Client struct {
	ID          int
	Name        string
	BDate       time.Time
	PhoneNumber string
	UpdatedAt   time.Time
}

func NewClient(name string, bDate time.Time, phoneNumber string, updatedAt time.Time) *Client {
	return &Client{
		Name:        name,
		BDate:       bDate,
		PhoneNumber: phoneNumber,
		UpdatedAt:   updatedAt,
	}
}

func (c *Client) SetID(id int) {
	c.ID = id
}

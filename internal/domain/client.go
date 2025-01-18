package domain

import (
	"time"
)

type Client struct {
	ID          int
	Name        string
	BDate       time.Time
	PhoneNumber string
}

func NewClient(name string, bDate time.Time, phoneNumber string) *Client {
	return &Client{
		Name:        name,
		BDate:       bDate,
		PhoneNumber: phoneNumber,
	}
}

func (c *Client) SetID(id int) {
	c.ID = id
}

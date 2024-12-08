package domain

type Client struct {
	ID          int
	Name        string
	BDate       string
	PhoneNumber string
}

func NewClient(name, bDate, phoneNumber string) *Client {
	return &Client{
		Name:        name,
		BDate:       bDate,
		PhoneNumber: phoneNumber,
	}
}

func (c *Client) SetID(id int) {
	c.ID = id
}

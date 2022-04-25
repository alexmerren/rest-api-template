package entities

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Contact struct {
	Age      int `json:"age"`
	ID       string
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
}

func NewContact(
	name string,
	age int,
	birthday string,
	address string,
	gender string,
) (*Contact, error) {
	return &Contact{
		ID:       NewUUID(),
		Name:     name,
		Age:      age,
		Birthday: birthday,
		Address:  address,
		Gender:   gender,
	}, nil
}

func (c Contact) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Age, validation.Required),
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.Birthday, validation.Required),
		validation.Field(&c.Address, validation.Required),
		validation.Field(&c.Gender, validation.Required),
	)
}

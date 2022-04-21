package entities

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Contact struct {
	Age      int
	ID       string
	Name     string
	Birthday string
	Address  string
	Gender   string
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

package entities

func MakeContact(
	name string,
	age int,
	birthday string,
	address string,
	Gender string,
) (*Contact, error) {
	id, err := UUID.New()
	if err != nil {
		return nil, err
	}

	return &Contact{
		ID:       id,
		Name:     name,
		Age:      age,
		Birthday: birthday,
		Address:  address,
		Gender:   gender,
	}, nil
}

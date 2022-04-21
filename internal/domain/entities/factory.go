package entities

import "math/rand"

const idLength = 10

func MakeContact(
	name string,
	age int,
	birthday string,
	address string,
	gender string,
) (Contact, error) {
	return Contact{
		ID:       generateNewID(),
		Name:     name,
		Age:      age,
		Birthday: birthday,
		Address:  address,
		Gender:   gender,
	}, nil
}

func generateNewID() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, idLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

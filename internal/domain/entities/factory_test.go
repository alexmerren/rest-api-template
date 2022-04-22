package entities_test

import (
	"rest-api-template/internal/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	test_name     = "name"
	test_age      = 1
	test_birthday = "someday"
	test_address  = "somewhere"
	test_gender   = "other"
)

func TestMakeContact(t *testing.T) {
	contact, err := entities.MakeContact(
		test_name,
		test_age,
		test_birthday,
		test_address,
		test_gender,
	)

	assert.Equal(t, 20, len(contact.ID))
	assert.Equal(t, test_name, contact.Name)
	assert.Equal(t, test_age, contact.Age)
	assert.Equal(t, test_birthday, contact.Birthday)
	assert.Equal(t, test_address, contact.Address)
	assert.Equal(t, test_gender, contact.Gender)
	assert.Nil(t, err)
}

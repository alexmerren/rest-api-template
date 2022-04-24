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

func TestNewContact(t *testing.T) {
	contact, err := entities.NewContact(
		test_name,
		test_age,
		test_birthday,
		test_address,
		test_gender,
	)

	assert.Equal(t, 36, len(contact.ID))
	assert.Equal(t, test_name, contact.Name)
	assert.Equal(t, test_age, contact.Age)
	assert.Equal(t, test_birthday, contact.Birthday)
	assert.Equal(t, test_address, contact.Address)
	assert.Equal(t, test_gender, contact.Gender)
	assert.Nil(t, err)
}

func TestContactValidation(t *testing.T) {
	var testCases = []struct {
		name           string
		contact        entities.Contact
		expectedErrMsg string
	}{
		{
			name:           "Empty Contact",
			contact:        entities.Contact{},
			expectedErrMsg: "address: cannot be blank; age: cannot be blank; birthday: cannot be blank; gender: cannot be blank; name: cannot be blank.",
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			err := tc.contact.Validate()
			assert.Equal(t, tc.expectedErrMsg, err.Error())
		})
	}
}

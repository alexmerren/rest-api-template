package entities_test

import (
	"rest-api-template/internal/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

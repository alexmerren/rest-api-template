package entities_test

import (
	"errors"
	"rest-api-template/internal/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNotFoundError(t *testing.T) {
	baseErr := errors.New("mock error")
	err := entities.NewNotFoundError("test", baseErr)

	assert.IsType(t, &entities.NotFoundError{}, err)
	assert.Equal(t, baseErr, err.Err)
	assert.Equal(t, "test", err.Error())
}

func TestNewInternalError(t *testing.T) {
	baseErr := errors.New("mock error")
	err := entities.NewInternalError("test", baseErr)

	assert.IsType(t, &entities.InternalError{}, err)
	assert.Equal(t, baseErr, err.Err)
	assert.Equal(t, "test", err.Error())
}

func TestNewBadRequestError(t *testing.T) {
	baseErr := errors.New("mock error")
	err := entities.NewBadRequestError("test", baseErr)

	assert.IsType(t, &entities.BadRequestError{}, err)
	assert.Equal(t, baseErr, err.Err)
	assert.Equal(t, "test", err.Error())
}

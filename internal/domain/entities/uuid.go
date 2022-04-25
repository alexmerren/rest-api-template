package entities

import gouuid "github.com/satori/go.uuid"

func NewUUID() string {
	return gouuid.NewV4().String()
}

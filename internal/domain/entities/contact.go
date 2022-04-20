package entities

type Contact struct {
	ID       uuid.UUID
	Age      int
	Name     string
	Birthday string
	Address  string
	Gender   string
}

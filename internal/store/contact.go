package store

type Contact struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone-number"`
}

func InsertContact(contact Contact) error {
	_, err := Exec(
		"INSERT INTO contacts (id, name, email, phonenumber) values (?, ?, ?, ?)",
		contact.ID, contact.Name, contact.Email, contact.PhoneNumber,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetContact(id string) (Contact, error) {
	row, err := Query("SELECT * FROM contacts WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func GetAllContacts() ([]Contact, error) {}

func UpdateContact(id string, contact Contact) error {}

func DeleteContact(id string) error {}

package store

type Contact struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
}

func (s *Store) InsertContact(contact *Contact) error {
	// Exec an insert record with contact information
	_, err := s.Exec(
		"INSERT INTO contact (name, email, phonenumber) values (?, ?, ?)",
		contact.Name, contact.Email, contact.PhoneNumber,
	)
	if err != nil {
		return err
	}

	// Get the ID of the inserted record and add that to the returning struct
	result := s.QueryRow("SELECT LAST_INSERT_ID()")
	err = result.Scan(&contact.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetContact(id string) (*Contact, error) {
	// Query the database for the relevant id, and scan that into the returning struct
	contact := &Contact{}
	result := s.QueryRow("SELECT * FROM contact WHERE id=?", id)
	err := result.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (s *Store) GetAllContacts() ([]*Contact, error) {
	rows, err := s.Query("SELECT * FROM contact")
	if err != nil {
		return nil, err
	}

	output := make([]*Contact, 0)
	for rows.Next() {
		contact := &Contact{}
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.PhoneNumber)
		if err != nil {
			return nil, err
		}
		output = append(output, contact)
	}
	return output, nil
}

func (s *Store) UpdateContact(contact *Contact) error {
	_, err := s.Exec(
		"UPDATE contact SET name=?, email=?, phonenumber=? WHERE id=?",
		contact.Name, contact.Email, contact.PhoneNumber, contact.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteContact(id string) error {
	_, err := s.Exec("DELETE FROM contact WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}

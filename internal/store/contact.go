package store

import "context"

type Contact struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
}

func (s *Store) InsertContact(context context.Context, contact *Contact) error {
	// Exec an insert record with contact information
	_, err := s.ExecContext(
		context, "INSERT INTO contact (name, email, phonenumber) values (?, ?, ?)",
		contact.Name, contact.Email, contact.PhoneNumber,
	)
	if err != nil {
		return err
	}

	// Get the ID of the inserted record and add that to the returning struct
	result := s.QueryRowContext(context, "SELECT LAST_INSERT_ID()")
	err = result.Scan(&contact.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetContact(context context.Context, id string) (*Contact, error) {
	// Query the database for the relevant id, and scan that into the returning struct
	contact := &Contact{}
	result := s.QueryRowContext(context, "SELECT * FROM contact WHERE id=?", id)
	err := result.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (s *Store) GetAllContacts(context context.Context) ([]*Contact, error) {
	rows, err := s.QueryContext(context, "SELECT * FROM contact")
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

func (s *Store) UpdateContact(context context.Context, contact *Contact) error {
	_, err := s.ExecContext(
		context, "UPDATE contact SET name=?, email=?, phonenumber=? WHERE id=?",
		contact.Name, contact.Email, contact.PhoneNumber, contact.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteContact(context context.Context, id string) error {
	_, err := s.ExecContext(context, "DELETE FROM contact WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}

package drivers

import (
	"errors"
	"log"
)

// AddUser adds a new user to the system.
func (td *TicketDriverBbolt) AddUser(user User) error {
	return errors.New("bbolt: Not implemented")
}

// GetUser returns a user struct for the selected user based on id.
func (td *TicketDriverBbolt) GetUser(id int) (User, error) {
	log.Println("bbolt: Get user")
	/*return User{ID: id}, nil*/
	return User{}, errors.New("bbolt: Not implemented")
}

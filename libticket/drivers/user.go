package drivers

import (
	"errors"
	"log"
)

func (td *TicketDriverBbolt) AddUser(user User) error {
	return errors.New("bbolt: Not implemented")
}

func (td *TicketDriverBbolt) GetUser(id int) (User, error) {
	log.Println("bbolt: Get user")
	/*return User{ID: id}, nil*/
	return User{}, errors.New("bbolt: Not implemented")
}

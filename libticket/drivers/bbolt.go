package drivers

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

var (
	Major int = 0
	Minor int = 1
)

type TicketDriverBbolt struct {
	DB *bolt.DB
}

func (td *TicketDriverBbolt) Connect() error {
	var err error
	log.Println("bbolt: Opening database file")
	td.DB, err = bolt.Open("antidote.db", 0600, nil)
	if err != nil {
		return err
	}
	log.Println("bbolt: Done opening database file")

	return nil
}

func (td *TicketDriverBbolt) Disconnect() error {
	log.Println("bbolt: Closing database file")
	td.DB.Close()

	return nil
}

func (td *TicketDriverBbolt) Info() string {
	return fmt.Sprintf("bbolt: driver v.%d.%d", Major, Minor)
}

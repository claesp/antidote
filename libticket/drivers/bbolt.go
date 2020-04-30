package drivers

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

var (
	// Major version of driver.
	Major = 0
	// Minor version of driver.
	Minor = 1
)

// TicketDriverBbolt is the default driver structure.
type TicketDriverBbolt struct {
	DB *bolt.DB
}

// Connect connects to the database file.
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

// Disconnect closes the database file.
func (td *TicketDriverBbolt) Disconnect() error {
	log.Println("bbolt: Closing database file")
	td.DB.Close()

	return nil
}

// Info displays the current driver information.
func (td *TicketDriverBbolt) Info() string {
	return fmt.Sprintf("bbolt: driver v.%d.%d", Major, Minor)
}

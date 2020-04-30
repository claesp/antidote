package libticket

import (
	"log"

	"github.com/claesp/antidote/libticket/drivers"
)

// TicketDB represents the current database (i.e. ticket system).
type TicketDB struct {
	CurrentDriverName string
	CurrentDriver     drivers.Driver
}

// Connect connects to the ticket system.
func (tdb *TicketDB) Connect() error {
	return tdb.CurrentDriver.Connect()
}

// Disconnect disconnects from the ticket system.
func (tdb *TicketDB) Disconnect() error {
	return tdb.CurrentDriver.Disconnect()
}

// GetUser returns a user from the ticket system, based on id.
func (tdb *TicketDB) GetUser(id int) (TicketUser, error) {
	var user TicketUser
	u, err := tdb.CurrentDriver.GetUser(id)
	if err != nil {
		log.Println(err)
		return user, err
	}
	user.ID = u.ID

	return user, nil
}

// Info returns information about the ticket system.
func (tdb *TicketDB) Info() string {
	return tdb.CurrentDriver.Info()
}

// Register registers a driver within the system.
func (tdb *TicketDB) Register(name string, driver drivers.Driver) {
	tdb.CurrentDriverName = name
	tdb.CurrentDriver = driver
}

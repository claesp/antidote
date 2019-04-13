package libticket

import (
	"log"

	"github.com/claesp/antidote/libticket/drivers"
)

type TicketDB struct {
	CurrentDriverName string
	CurrentDriver     drivers.Driver
}

func (tdb *TicketDB) Connect() error {
	return tdb.CurrentDriver.Connect()
}

func (tdb *TicketDB) Disconnect() error {
	return tdb.CurrentDriver.Disconnect()
}

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

func (tdb *TicketDB) Register(name string, driver drivers.Driver) {
	tdb.CurrentDriverName = name
	tdb.CurrentDriver = driver
}

package libticket

import "github.com/claesp/antidote/libticket/drivers"

type TicketDB struct {
	CurrentDriverName string
	CurrentDriver     drivers.Driver
}

func (tdb *TicketDB) Connect() (bool, error) {
	return tdb.CurrentDriver.Connect()
}

func (tdb *TicketDB) Register(name string, driver drivers.Driver) {
	tdb.CurrentDriverName = name
	tdb.CurrentDriver = driver
}

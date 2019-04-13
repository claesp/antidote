package drivers

import (
	"errors"
	"fmt"
)

var (
	Major int = 0
	Minor int = 1
)

type TicketDriverBoltDB struct {
}

func (td *TicketDriverBoltDB) Connect() error {
	return errors.New("Not implemented")
}

func (td *TicketDriverBoltDB) Info() string {
	return fmt.Sprintf("BoltDB driver v.%d.%d", Major, Minor)
}

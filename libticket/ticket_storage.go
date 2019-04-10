package libticket

type TicketStorage interface {
	Init() (bool, error)
}

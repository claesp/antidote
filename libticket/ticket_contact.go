package libticket

type TicketContact struct {
	ID    int
	Email string
	Name  string
	Phone string
}

func (tc TicketContact) HasEmail() bool {
	return len(tc.Email) > 0
}

func (tc TicketContact) HasPhone() bool {
	return len(tc.Phone) > 0
}

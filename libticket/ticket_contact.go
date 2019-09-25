package libticket

// TicketContact is a structure for the representation of the contact on a
// ticket.
type TicketContact struct {
	ID    int
	Email string
	Name  string
	Phone string
}

// HasEmail returns true if the ticket contact has a e-mail address set.
func (tc TicketContact) HasEmail() bool {
	return len(tc.Email) > 0
}

// HasPhone returns true if the ticket contact has a phone number set.
func (tc TicketContact) HasPhone() bool {
	return len(tc.Phone) > 0
}

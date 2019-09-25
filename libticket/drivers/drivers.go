package drivers

// User is a user of the system.
type User struct {
	ID int
}

// Driver interface for a ticket driver.
type Driver interface {
	AddUser(User) error
	Connect() error
	Disconnect() error
	GetUser(id int) (User, error)
	Info() string
}

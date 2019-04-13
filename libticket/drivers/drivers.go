package drivers

type User struct {
	ID int
}

type Driver interface {
	Connect() error
	Disconnect() error
	GetUser(id int) (User, error)
	Info() string
}

package drivers

type User struct {
	ID int
}

type Driver interface {
	AddUser(User) error
	Connect() error
	Disconnect() error
	GetUser(id int) (User, error)
	Info() string
}

package drivers

type Driver interface {
	Connect() (bool, error)
	Info() string
}

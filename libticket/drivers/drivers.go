package drivers

type Driver interface {
	Connect() error
	Info() string
}

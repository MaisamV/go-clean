package delivery

type Server interface {
	Init() error
	Serv() error
}

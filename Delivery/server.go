package Delivery

type Server interface {
	Init() error
	Serv() error
}

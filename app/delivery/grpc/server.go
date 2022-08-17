package grpc

import (
	"GoCleanMicroservice/abstract/domain"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("grpc_port", 50051, "The grpc server port")
)

type Server struct {
	Engine      *grpc.Server
	Interactors *domain.InteractorPackage
}

func (m *Server) Init() error {
	m.Engine = grpc.NewServer()
	m.setInteractors()
	m.addServices()
	return nil
}

func (m *Server) Serv() error {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", flag.Lookup("grpc_port").Value, err)
	}
	log.Printf("server listening at %v", lis.Addr())
	if err = m.Engine.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return err
}

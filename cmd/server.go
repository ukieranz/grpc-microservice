package cmd

import (
	"fmt"
	"log"
	"modules/v2/app/service"
	"modules/v2/pkg/config"
	"modules/v2/pkg/generated/v1/greeter"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	address string
}

func NewServer() *Server {
	hostAddr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	return &Server{
		address: hostAddr,
	}
}

func (s *Server) GrpcListen() error {
	_listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	_grpc := grpc.NewServer()
	greeter.RegisterGreeterServer(_grpc, &service.Greeter{})
	log.Printf("grpc listening on %s\n", _listener.Addr())
	return _grpc.Serve(_listener)
}

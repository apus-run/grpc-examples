package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	v1 "github.com/apus-run/grpc-examples/user-service/api/users/v1"
	"github.com/apus-run/grpc-examples/user-service/service"
)

func registerServices(s *grpc.Server) {
	v1.RegisterUsersServer(s, service.NewUserService())
}

func startServer(s *grpc.Server, l net.Listener) error {
	return s.Serve(l)
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":50051"
	}

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	registerServices(s)
	log.Fatal(startServer(s, lis))
}

package main

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	v1 "github.com/apus-run/grpc-examples/user-service/api/users/v1"
)

type dummyUserService struct {
	v1.UnimplementedUsersServer
}

func (s *dummyUserService) GetUser(ctx context.Context, in *v1.UserGetRequest) (*v1.UserGetReply, error) {
	u := v1.User{
		Id:        "user-123-a",
		FirstName: "jane",
		LastName:  "doe",
		Age:       36,
	}
	return &v1.UserGetReply{User: &u}, nil
}

func startTestGrpcServer() (*grpc.Server, *bufconn.Listener) {
	l := bufconn.Listen(10)
	s := grpc.NewServer()
	v1.RegisterUsersServer(s, &dummyUserService{})
	go func() {
		err := s.Serve(l)
		if err != nil {
			log.Fatal(err)
		}
	}()
	return s, l
}

func TestGetUser(t *testing.T) {
	s, l := startTestGrpcServer()
	defer s.GracefulStop()

	bufconnDialer := func(
		ctx context.Context, addr string,
	) (net.Conn, error) {
		return l.Dial()
	}

	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(bufconnDialer),
	)
	if err != nil {
		t.Fatal(err)
	}

	c := getUserServiceClient(conn)
	result, err := getUser(
		c,
		&v1.UserGetRequest{Email: "moocss@163.com"},
	)
	if err != nil {
		t.Fatal(err)
	}

	if result.User.FirstName != "jane" ||
		result.User.LastName != "doe" {
		t.Fatalf(
			"Expected: jane doe, Got: %s %s",
			result.User.FirstName,
			result.User.LastName,
		)
	}
}

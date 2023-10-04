package service

import (
	"context"
	"errors"
	"log"
	"strings"

	v1 "github.com/apus-run/grpc-examples/user-service/api/users/v1"
)

type userService struct {
	v1.UnimplementedUsersServer
}

func NewUserService() *userService {
	return &userService{}
}

func (s *userService) GetUser(ctx context.Context, in *v1.UserGetRequest) (*v1.UserGetReply, error) {
	log.Printf(
		"Received request for user with Email: %s Id: %s\n",
		in.Email,
		in.Id,
	)
	components := strings.Split(in.Email, "@")
	if len(components) != 2 {
		return nil, errors.New("invalid email address")
	}
	u := v1.User{
		Id:        in.Id,
		FirstName: components[0],
		LastName:  components[1],
		Age:       36,
	}
	return &v1.UserGetReply{User: &u}, nil
}

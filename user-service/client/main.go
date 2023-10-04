package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	v1 "github.com/apus-run/grpc-examples/user-service/api/users/v1"
)

func setupGrpcConn(addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(
		context.Background(),
		addr,
		// grpc.WithInsecure(), // 已弃用, 使用下面的方式代替
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
}

func getUserServiceClient(conn *grpc.ClientConn) v1.UsersClient {
	return v1.NewUsersClient(conn)
}

func getUser(client v1.UsersClient, u *v1.UserGetRequest) (*v1.UserGetReply, error) {
	return client.GetUser(context.Background(), u)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal(
			"Must specify a gRPC server address",
		)
	}
	conn, err := setupGrpcConn(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := getUserServiceClient(conn)

	result, err := getUser(
		c,
		&v1.UserGetRequest{Email: "jane@doe.com"},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(
		os.Stdout, "User: %s %s\n",
		result.User.FirstName,
		result.User.LastName,
	)
}

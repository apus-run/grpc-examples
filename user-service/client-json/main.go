package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	v1 "github.com/apus-run/grpc-examples/user-service/api/users/v1"
)

func setupGrpcConn(addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(
		context.Background(),
		addr,
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

func createUserRequest(jsonQuery string) (*v1.UserGetRequest, error) {
	u := v1.UserGetRequest{}
	input := []byte(jsonQuery)
	return &u, protojson.Unmarshal(input, &u)
}

func getUserResponseJson(result *v1.UserGetReply) ([]byte, error) {
	return protojson.Marshal(result)
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal(
			"Must specify a gRPC server address and search query",
		)
	}
	serverAddr := os.Args[1]

	u, err := createUserRequest(os.Args[2])
	if err != nil {
		log.Fatalf("Bad user input: %v", err)
	}

	conn, err := setupGrpcConn(serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := getUserServiceClient(conn)

	result, err := getUser(
		c,
		u,
	)
	if err != nil {
		log.Fatal(err)
	}

	data, err := getUserResponseJson(result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(
		os.Stdout, string(data),
	)
}

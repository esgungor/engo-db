package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/esgungor/engo-db/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func parseUserCommand(arg []string) (*api.CommandRequest, error) {
	req := &api.CommandRequest{
		Operation: arg[1],
		Key:       arg[2],
	}
	if arg[1] == "INSERT" || arg[1] == "UPDATE" {
		req.Value = arg[3]
	}
	return req, nil
}
func sendCommand(client api.CommandClient, command *api.CommandRequest) (*api.CommandReply, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	resp, err := client.Command(ctx, command)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func main() {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	path := "localhost:50000"
	conn, err := grpc.Dial(path, opts)
	if err != nil {
		panic(err)

	}
	client := api.NewCommandClient(conn)
	input := os.Args
	conv, err := parseUserCommand(input)
	if err != nil {
		panic(err)
	}
	resp, err := sendCommand(client, conv)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}

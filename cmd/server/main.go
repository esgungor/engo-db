package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/esgungor/engo-db/api"
	"github.com/esgungor/engo-db/pkg/store"
	"github.com/esgungor/engo-db/pkg/wal"
	"google.golang.org/grpc"
)

type commandServer struct {
	keyValue *store.KeyValueStore
	api.UnimplementedCommandServer
}

func (c *commandServer) Command(ctx context.Context, in *api.CommandRequest) (*api.CommandReply, error) {
	if in.Operation == "GET" {
		val, err := c.keyValue.Get(in.Key)
		if err != nil {
			return nil, err
		}
		return &api.CommandReply{
			Key:   in.Key,
			Value: val,
		}, nil
	}
	if err := c.keyValue.Insert(store.Entry{
		Timestamp: time.Now(),
		Operation: in.Operation,
		Key:       in.Key,
		Value:     in.Value,
	}); err != nil {
		return nil, err
	}
	return &api.CommandReply{
		Message: "OK",
	}, nil
}

func newCommandServer() api.CommandServer {
	wal, err := wal.NewWAL("wal.log")
	if err != nil {
		panic(err)
	}
	s := store.NewKeyValueStore(wal)
	s.Init()
	return &commandServer{
		keyValue: &s,
	}
}

func main() {
	fmt.Println("OK")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50000))
	if err != nil {
		log.Fatalf("starting server failed %v", err)
	}
	server := grpc.NewServer()
	api.RegisterCommandServer(server, newCommandServer())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("error when starting server: %v", err)
	}
}

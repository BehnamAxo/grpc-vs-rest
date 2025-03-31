package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync/atomic"
	"time"

	"grpc/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type server struct {
	userpb.UnimplementedPaymentServiceServer
}

var requestCounter uint64

func (s *server) ProcessUser(ctx context.Context, user *userpb.User) (*userpb.UserAck, error) {
	count := atomic.AddUint64(&requestCounter, 1)

	log.Printf("[%d] Received: %s | Interests: %d | History Events: %d | Bio Size: %d bytes\n",
		count, user.Name, len(user.Interests), len(user.History), len(user.Bio))

	return &userpb.UserAck{
		Message: "Payment processed for " + user.Name,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("❌ Failed to listen:", err)
	}

	kaParams := keepalive.ServerParameters{
		MaxConnectionIdle:     5 * time.Minute,
		MaxConnectionAge:      10 * time.Minute,
		MaxConnectionAgeGrace: 5 * time.Minute,
		Time:                  2 * time.Minute,
		Timeout:               20 * time.Second,
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(kaParams),
	)

	userpb.RegisterPaymentServiceServer(grpcServer, &server{})

	fmt.Println("gRPC Payment service running on port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}

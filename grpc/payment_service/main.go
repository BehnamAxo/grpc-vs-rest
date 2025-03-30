package main

import (
	"context"
	"fmt"
	"grpc/generated/userpb"
	"log"
	"net"
	"sync/atomic"

	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedPaymentServiceServer
}

var requestCounter uint64

func (s *server) ProcessUser(ctx context.Context, user *userpb.User) (*userpb.UserAck, error) {
	count := atomic.AddUint64(&requestCounter, 1)

	log.Printf("[#%d] Payment received for user: %s\n", count, user.Name)
	return &userpb.UserAck{Message: "Payment processed for " + user.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterPaymentServiceServer(grpcServer, &server{})

	fmt.Println("gRPC Payment service running on port 50051!")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

// ghz --insecure --proto ./proto/user.proto --call userpb.PaymentService.ProcessUser -d "{\"name\":\"Sir Laughsalot McGiggles\",\"age\":420,\"email\":\"funny.bone@laughterverse.io\",\"phone\":\"+1-800-GIGGLEZ\"}" -c 100 --duration 20s localhost:50051

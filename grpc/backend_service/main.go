package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"grpc/generated/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	client := userpb.NewPaymentServiceClient(conn)

	user := &userpb.User{
		Name:  "Sir Laughsalot McGiggles",
		Age:   420,
		Email: "funny.bone@laughterverse.io",
		Phone: "+1-800-GIGGLEZ",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ProcessUser(ctx, user)
	if err != nil {
		log.Fatalf("Error calling payment service: %v", err)
	}

	fmt.Println("✅ Response from payment service:", res.Message)
}

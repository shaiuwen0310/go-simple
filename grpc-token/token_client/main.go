package main

import (
	"context"
	"log"

	pb "grpc-token/auth"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTokenServiceClient(conn)

	// Generate a token
	password := "password123"
	genResp, err := client.GenToken(context.Background(), &pb.PasswordRequest{Password: password})
	if err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}
	log.Println("Token:", genResp.Token)

	// Parse the token
	parseResp, err := client.ParseToken(context.Background(), &pb.TokenRequest{Token: genResp.Token})
	if err != nil {
		log.Fatalf("Failed to parse token: %v", err)
	}
	log.Println("Password:", parseResp.Password)
}

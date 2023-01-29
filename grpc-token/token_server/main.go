package main

import (
	"context"
	"errors"
	"log"
	"net"
	"time"

	pb "grpc-token/auth"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
)

type MyClaims struct {
	Password string `json:"password"`
	jwt.RegisteredClaims
}

const TokenExpireDuration = time.Second * 60

var MySecret = []byte("secret-key")

type server struct {
	pb.UnimplementedTokenServiceServer
}

func (s *server) GenToken(ctx context.Context, in *pb.PasswordRequest) (*pb.TokenResponse, error) {
	c := MyClaims{
		in.Password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "judy",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedToken, err := token.SignedString(MySecret)
	if err != nil {
		return nil, err
	}
	return &pb.TokenResponse{Token: signedToken}, nil
}

func (s *server) ParseToken(ctx context.Context, in *pb.TokenRequest) (*pb.ClaimsResponse, error) {
	token, err := jwt.ParseWithClaims(in.Token, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return &pb.ClaimsResponse{Password: claims.Password}, nil
	}
	return nil, errors.New("invalid token")
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTokenServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

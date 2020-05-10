package client

import (
	"github.com/maei/shared_utils_go/logger"
	"google.golang.org/grpc"
	"log"
	"os"
)

var GRPCClient grpcClientInterface = &grpcClient{}

type grpcClientInterface interface {
	SetClient() (*grpc.ClientConn, error)
}

type grpcClient struct{}

func (*grpcClient) SetClient() (*grpc.ClientConn, error) {
	logger.Info("starting gRPC Client")
	conn, err := grpc.Dial(os.Getenv("SERVER_HOST"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
		return nil, err
	}

	return conn, nil
}

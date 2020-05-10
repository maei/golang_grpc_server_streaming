package server

import (
	"fmt"
	"github.com/maei/golang_grpc_server_streaming/grpc_server_stream/src/domain/greetpb"
	"github.com/maei/golang_grpc_server_streaming/grpc_server_stream/src/services"
	"github.com/maei/shared_utils_go/logger"
	"google.golang.org/grpc"
	"net"
	"os"
	"time"
)

type server struct{}

var (
	s = grpc.NewServer()
)

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	greet := services.GreetService.Greet(firstName, lastName)

	logger.Info("streaming-started")
	for i := 0; i < 10; i++ {
		res := &greetpb.GreetManyTimesResponse{
			Result: fmt.Sprintf("%s and the counter is %v", greet, i),
		}
		stream.Send(res)
		time.Sleep(1 * time.Second)
	}
	logger.Info("streaming-ended")
	return nil
}

func StartGRPCServer() {
	logger.Info("gRPC greet-streaming started")

	lis, err := net.Listen("tcp", os.Getenv("SERVER_PORT"))
	if err != nil {
		logger.Error("error while listening gRPC Server", err)
	}
	greetpb.RegisterGreetServiceServer(s, &server{})

	errServer := s.Serve(lis)
	if errServer != nil {
		logger.Error("error while serve gRPC Server", errServer)
	}
}

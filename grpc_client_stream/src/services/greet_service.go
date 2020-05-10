package services

import (
	"context"
	"fmt"
	"github.com/maei/golang_grpc_server_streaming/grpc_client_stream/src/client"
	"github.com/maei/golang_grpc_server_streaming/grpc_client_stream/src/domain/greetpb"
	"github.com/maei/shared_utils_go/logger"
	"io"
)

var GreetService greetServiceInterface = &greetService{}

type greetServiceInterface interface {
	Greet()
}

type greetService struct{}

func (*greetService) Greet() {
	conn, err := client.GRPCClient.SetClient()
	if err != nil {
		logger.Error("error while connecting grpc client", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Matthias",
			LastName:  "Eiletz",
		},
	}

	resStream, greetErr := c.GreetManyTimes(context.Background(), req)
	if greetErr != nil {
		logger.Error("error while getting stream", greetErr)
	}

	for {
		msg, streamErr := resStream.Recv()
		if streamErr == io.EOF {
			// End of stream
			break
		}
		if streamErr != nil {
			logger.Error("error while getting messages from stream", streamErr)
		}
		fmt.Printf("Response from steam: %v\n", msg.GetResult())
	}
}

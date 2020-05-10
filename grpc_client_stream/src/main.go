package main

import (
	"github.com/maei/golang_grpc_server_streaming/grpc_client_stream/src/services"
	"github.com/maei/shared_utils_go/logger"
)

func main() {
	logger.Info("grpc-streaming client started")
	services.GreetService.Greet()
}

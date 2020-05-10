package main

import (
	"github.com/maei/golang_grpc_server_streaming/grpc_client_stream/src/services"
	"github.com/maei/shared_utils_go/logger"
	"time"
)

func main() {
	logger.Info("grpc-streaming client started")
	time.Sleep(5 * time.Second)
	services.GreetService.Greet()
}

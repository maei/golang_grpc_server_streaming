FROM golang:latest as builder

ENV GO111MODULE=auto
ENV SERVER_HOST=grpc-server:50051

WORKDIR /go/src/github.com/maei/golang_grpc_server_streaming/grpc_client_stream

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o grpc-stream-client /go/src/github.com/maei/golang_grpc_server_streaming/grpc_client_stream/src/main.go

FROM alpine
COPY --from=builder /go/src/github.com/maei/golang_grpc_server_streaming/grpc_client_stream/grpc-stream-client /
CMD ["./grpc-stream-client"]
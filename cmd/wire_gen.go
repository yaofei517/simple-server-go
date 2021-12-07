// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	grpc2 "google.golang.org/grpc"
	http2 "net/http"
	"simple-server-go/api/grpc"
	"simple-server-go/api/http"
	"simple-server-go/api/http/middle"
	"simple-server-go/internal/endpoint"
)

// Injectors from wire.go:

// initApp init endpoints
func InitServer() (*Server, error) {
	endpoints := endpoint.NewEndpoints()
	middleware, err := middle.NewMiddleware()
	if err != nil {
		return nil, err
	}
	server := http.NewHttp(endpoints, middleware)
	grpcServer, err := grpc.NewGrpc()
	if err != nil {
		return nil, err
	}
	cmdServer := NewServer(server, grpcServer)
	return cmdServer, nil
}

// wire.go:

type Server struct {
	Http *http2.Server
	Grpc *grpc2.Server
}

func NewServer(http3 *http2.Server, grpc3 *grpc2.Server) *Server {
	return &Server{
		Grpc: grpc3,
		Http: http3,
	}
}

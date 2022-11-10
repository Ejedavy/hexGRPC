package rpc

import (
	"hex/internal/adapters/framework/left/grpc/pb"
	"hex/internal/ports"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api ports.APIPort
}

func NewGRPCAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (adapter Adapter) Run() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Could not listen at port 9000: %v", err)
	}

	new_grpc := grpc.NewServer()
	service := adapter
	pb.RegisterArithmeticServiceServer(new_grpc, service)
	reflection.Register(new_grpc)

	if err := new_grpc.Serve(listen); err != nil {
		log.Fatalf("Error starting the server %v", err)
	}
}

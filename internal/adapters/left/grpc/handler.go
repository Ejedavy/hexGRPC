package rpc

import (
	"context"
	"github.com/ejedavy/hexGRPC/internal/adapters/left/grpc/proto"
	"github.com/ejedavy/hexGRPC/internal/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type handler struct {
	application *app.ArithmeticApp
	proto.UnimplementedArithmeticServiceServer
}

func NewGRPCAdapter(ap *app.ArithmeticApp) *handler {
	return &handler{application: ap}
}

func (h *handler) Run() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Could not listen at port 9000: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterArithmeticServiceServer(grpcServer, h)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Error starting the server %v", err)
	}
}

func (h *handler) GetAddition(ctx context.Context, req *proto.Operands) (*proto.Answer, error) {
	var err error
	var ans *proto.Answer

	if req.A == 0 || req.B == 0 {
		return ans, status.Error(codes.InvalidArgument, "You provided invalid arguments")
	}

	answer, err := h.application.GetAddition(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	ans = &proto.Answer{Value: answer}
	return ans, nil
}

func (h *handler) GetSubtraction(ctx context.Context, req *proto.Operands) (*proto.Answer, error) {
	var err error
	var ans *proto.Answer

	if req.A == 0 || req.B == 0 {
		return ans, status.Error(codes.InvalidArgument, "You provided invalid arguments")
	}

	answer, err := h.application.GetSubtraction(req.A, req.B)

	if err != nil {
		return &proto.Answer{Value: answer}, status.Error(codes.Internal, "Unexpected Error")
	}

	ans = &proto.Answer{Value: answer}
	return ans, nil
}

func (h *handler) GetMultiplication(ctx context.Context, req *proto.Operands) (*proto.Answer, error) {
	var err error
	var ans *proto.Answer

	if req.A == 0 || req.B == 0 {
		return ans, status.Error(codes.InvalidArgument, "You provided invalid arguments")
	}

	answer, err := h.application.GetMultiplication(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	ans = &proto.Answer{Value: answer}
	return ans, nil
}

func (h *handler) GetDivision(ctx context.Context, req *proto.Operands) (*proto.Answer, error) {
	var err error
	var ans *proto.Answer

	if req.A == 0 || req.B == 0 {
		return ans, status.Error(codes.InvalidArgument, "You provided invalid arguments")
	}

	answer, err := h.application.GetDivision(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	ans = &proto.Answer{Value: answer}
	return ans, nil
}

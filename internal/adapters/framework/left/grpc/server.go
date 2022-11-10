package rpc

import (
	"context"
	"hex/internal/adapters/framework/left/grpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (adapter Adapter) GetAddition(ctx context.Context, req *pb.Operands) (*pb.Answer, error) {
	var err error
	var ans *pb.Answer

	if req.A == 0 || req.B == 0 {
		return ans, status.Error(codes.InvalidArgument, "You provided invalid arguments")
	}

	answer, err := adapter.api.GetAddition(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	ans = &pb.Answer{Value: answer}
	return ans, nil
}

func (adapter Adapter) GetSubstraction(ctx context.Context, req *pb.Operands) (*pb.Answer, error) {
	var err error
	var ans *pb.Answer

	if req.A == 0 || req.B == 0 {
		return ans, status.Error(codes.InvalidArgument, "You provided invalid arguments")
	}

	answer, err := adapter.api.GetSubstraction(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	ans = &pb.Answer{Value: answer}
	return ans, nil
}

func (adapter Adapter) GetMultiplication(ctx context.Context, req *pb.Operands) (*pb.Answer, error) {
	var err error
	var ans *pb.Answer

	if req.A == 0 || req.B == 0 {
		return ans, status.Error(codes.InvalidArgument, "You provided invalid arguments")
	}

	answer, err := adapter.api.GetMultiplication(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	ans = &pb.Answer{Value: answer}
	return ans, nil
}

func (adapter Adapter) GetDivision(ctx context.Context, req *pb.Operands) (*pb.Answer, error) {
	var err error
	var ans *pb.Answer

	if req.A == 0 || req.B == 0 {
		return ans, status.Error(codes.InvalidArgument, "You provided invalid arguments")
	}

	answer, err := adapter.api.GetDivision(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "Unexpected Error")
	}

	ans = &pb.Answer{Value: answer}
	return ans, nil
}

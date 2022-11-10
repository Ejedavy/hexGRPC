package ports

import (
	"context"
	"hex/internal/adapters/framework/left/grpc/pb"
)

type GRPCPort interface {
	Run()
	GetAddition(context.Context, *pb.Operands) (*pb.Answer, error)
	GetSubstraction(context.Context, *pb.Operands) (*pb.Answer, error)
	GetMultiplication(context.Context, *pb.Operands) (*pb.Answer, error)
	GetDivision(context.Context, *pb.Operands) (*pb.Answer, error)
}

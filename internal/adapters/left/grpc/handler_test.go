package rpc

import (
	"context"
	"github.com/ejedavy/hexGRPC/internal/adapters/left/grpc/proto"
	"github.com/ejedavy/hexGRPC/internal/adapters/right/db"
	"github.com/ejedavy/hexGRPC/internal/app"
	"github.com/ejedavy/hexGRPC/internal/ports"
	"log"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var database ports.DBPort
	var err error
	lis = bufconn.Listen(bufSize)

	driverName := os.Getenv("DB_DRIVER")
	dataSourceName := os.Getenv("DS_NAME")
	database, err = db.NewAdapter(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("Cannot start up the database %v", err)
	}

	defer database.CloseConnection()
	application := app.New(database)

	GRPC := grpc.NewServer()
	server := NewGRPCAdapter(application)
	proto.RegisterArithmeticServiceServer(GRPC, server)

	go func() {
		if err := GRPC.Serve(lis); err != nil {
			log.Fatalf("test server start error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	return conn
}

func TestGetAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := proto.NewArithmeticServiceClient(conn)

	params := &proto.Operands{
		A: 1,
		B: 1,
	}

	answer, err := client.GetAddition(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer.Value, int32(2))
}

func TestGetMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := proto.NewArithmeticServiceClient(conn)

	params := &proto.Operands{
		A: 1,
		B: 1,
	}

	ans, err := client.GetMultiplication(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}
	require.Equal(t, ans.Value, int32(1))

}

func TestGetDivision(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := proto.NewArithmeticServiceClient(conn)

	params := &proto.Operands{
		A: 1,
		B: 1,
	}

	ans, err := client.GetDivision(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}
	require.Equal(t, ans.Value, int32(1))

}

func TestGetSubtraction(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := proto.NewArithmeticServiceClient(conn)

	params := &proto.Operands{
		A: 1,
		B: 1,
	}

	ans, err := client.GetSubtraction(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}
	require.Equal(t, ans.Value, int32(0))

}

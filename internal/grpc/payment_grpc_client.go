package grpc

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/milo1150/cart-demo-proto/pkg/payment"
)

func ConnectToPaymentGRPCServer(log *zap.Logger) *grpc.ClientConn {
	endpoint := os.Getenv("GRPC_PAYMENT_ENDPOINT")

	conn, err := grpc.NewClient(
		endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Error(fmt.Sprintf("Failed to connect ShopProduct grpc server: %v", err))
	}

	return conn
}

func GetPayment(ctx context.Context, conn *grpc.ClientConn, paymentOrderId uint) (*pb.GetPaymentResponse, error) {
	client := pb.NewPaymentServiceClient(conn)

	newCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req := &pb.GetPaymentRequest{PaymentOrderId: uint64(paymentOrderId)}
	res, err := client.GetPayment(newCtx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

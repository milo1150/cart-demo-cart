package grpc

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/milo1150/cart-demo-proto/pkg/shop_product"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectToShopProductGRPCServer(log *zap.Logger) *grpc.ClientConn {
	endpoint := os.Getenv("GRPC_SHOP_PRODUCT_ENDPOINT")

	conn, err := grpc.NewClient(
		endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Error(fmt.Sprintf("Failed to connect ShopProduct grpc server: %v", err))
	}

	return conn
}

func ProductExists(ctx context.Context, conn *grpc.ClientConn, productId uint) (bool, error) {
	client := pb.NewShopProductServiceClient(conn)

	newCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel() // Use the provided context with a timeout to prevent infinite blocking

	req := &pb.CheckProductRequest{ProductId: uint64(productId)}
	res, err := client.ProductExists(newCtx, req)
	if err != nil {
		return false, err
	}

	return res.IsExists, nil
}

func GetProduct(ctx context.Context, conn *grpc.ClientConn, productId uint) (*pb.GetProductResponse, error) {
	client := pb.NewShopProductServiceClient(conn)

	newCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req := &pb.GetProductRequest{ProductId: uint64(productId)}
	res, err := client.GetProduct(newCtx, req)
	if err != nil {
		return nil, err
	}

	return res, err
}

func GetProducts(ctx context.Context, conn *grpc.ClientConn, productIds []uint64) (*pb.GetProductsResponse, error) {
	client := pb.NewShopProductServiceClient(conn)

	newCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req := &pb.GetProductsRequest{ProductIds: productIds}
	res, err := client.GetProducts(newCtx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

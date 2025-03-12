package grpc

import (
	"context"
	"log"
	"time"

	pb "github.com/milo1150/cart-demo-proto/pkg/shop_product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectToShopProductGRPCServer() *grpc.ClientConn {
	conn, err := grpc.NewClient(
		"demo-shop-product-service-app-1:50051", // TODO: do not hard code
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Printf("Failed to connect ShopProduct grpc server: %v", err)
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

package grpc

import (
	"context"
	"fmt"
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

func GetProduct(conn *grpc.ClientConn) {
	// Creates a gRPC client instance
	client := pb.NewShopProductServiceClient(conn)

	// Adds a timeout to the request (prevents infinite wait)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call GetProduct
	req := &pb.GetProductRequest{ProductId: 1}
	res, err := client.GetProduct(ctx, req)
	if err != nil {
		log.Printf("Error calling GetProduct: %v", err)
		return
	}

	// Print the response
	fmt.Printf("Product: ID=%d, Name=%s, Price=%.2f, Stock=%d, ShopID=%d\n",
		res.Id, res.ProductName, res.Price, res.Stock, res.ShopId)
}

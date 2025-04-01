package main

import (
	"cart-service/internal/database"
	"cart-service/internal/grpc"
	"cart-service/internal/loader"
	"cart-service/internal/middlewares"
	"cart-service/internal/nats"
	"cart-service/internal/routes"
	"cart-service/internal/types"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load ENV
	loader.LoadEnv()

	// NATS
	nc := nats.ConnectNATS()
	defer nc.Close()

	// NATS JetStream
	js := nats.ConnectJetStream(nc)

	// Database handler
	db := database.ConnectDatabase()
	database.RunAutoMigrate(db)

	// Initialize Zap Logger
	logger := middlewares.InitializeZapLogger()

	// Connect to gRPC Servers
	grpcShopProductClientConn := grpc.ConnectToShopProductGRPCServer(logger)
	grpcPaymentClientConn := grpc.ConnectToPaymentGRPCServer(logger)

	// Global state
	appState := &types.AppState{
		DB:                        db,
		NATS:                      nc,
		JS:                        js,
		Log:                       logger,
		GrpcShopProductClientConn: grpcShopProductClientConn,
		GrpcPaymentClientConn:     grpcPaymentClientConn,
	}

	// Creates an instance of Echo.
	e := echo.New()

	// Middlewares
	middlewares.RegisterMiddlewares(e)

	// Init Route
	routes.RegisterAppRoutes(e, appState)

	// Run NATS services
	go nats.SubscribeCreateUserEvent(nc, logger, db)
	go nats.PublishCreateCheckoutEvent(js, logger)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}

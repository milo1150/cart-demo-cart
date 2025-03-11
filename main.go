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

	// TODO: try JetStream

	// Database handler
	db := database.ConnectDatabase()
	database.RunAutoMigrate(db)

	// Initialize Zap Logger
	logger := middlewares.InitializeZapLogger()

	// Global state
	appState := &types.AppState{
		DB:   db,
		NATS: nc,
		Log:  logger,
	}

	// Creates an instance of Echo.
	e := echo.New()

	// Middlewares
	middlewares.RegisterMiddlewares(e)

	// Init Route
	routes.RegisterAppRoutes(e, appState)

	// Init NATS Pub/Sub
	nats.SubscribeToUserService(nc, logger, db)

	// Connect to gRPC Servers
	conn := grpc.ConnectToShopProductGRPCServer()

	// TODO: remove
	grpc.GetProduct(conn)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}

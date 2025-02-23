package main

import (
	"cart-service/internal/database"
	"cart-service/internal/loader"
	"cart-service/internal/middlewares"
	"cart-service/internal/routes"
	"cart-service/internal/types"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load ENV
	loader.LoadEnv()

	// Database handler
	db := database.ConnectDatabase()
	database.RunAutoMigrate(db)

	// Global state
	appState := &types.AppState{
		DB: db,
	}

	// Creates an instance of Echo.
	e := echo.New()

	// Middlewares
	middlewares.RegisterMiddlewares(e)

	// Init Route
	routes.RegisterAppRoutes(e, appState)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}

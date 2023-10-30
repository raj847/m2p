package main

import (
	"context"
	"m2p/repository"
	"m2p/service"
	"m2p/utils"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := utils.ConnectDB()
	if err != nil {
		panic(err)
	}
	mongoDB := utils.ConnectMongo(context.Background())
	defer utils.CloseMongo(context.Background())

	trxRepo := repository.NewTrxRepository(db, mongoDB)
	trxServ := service.NewTrxService(trxRepo)

	// Routes
	e.POST("/v1/register", trxServ.Create)

	// Start server
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}

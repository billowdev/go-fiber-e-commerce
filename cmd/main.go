package main

import (
	"fmt"
	"log"


	"github.com/billowdev/go-fiber-e-commerce/internal/adapters/app"
	"github.com/billowdev/go-fiber-e-commerce/internal/adapters/database"
	"github.com/billowdev/go-fiber-e-commerce/pkg/configs"
)

func main() {
	params := configs.NewFiberHttpServiceParams()
	fiberHTTP := configs.NewFiberHTTPService(params)
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal("Failed to start Database:", err)
	}

	app.AppContainer(fiberHTTP, db)
	portString := fmt.Sprintf(":%v", params.Port)

	err = fiberHTTP.Listen(portString)

	if err != nil {
		log.Fatal("Failed to start golang Fiber server:", err)
	}

}

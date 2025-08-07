package main

import (
	"cobaaja/contact-management/config"
	"cobaaja/contact-management/route"
	"fmt"
	"log"
	"os"
)

func main() {
	// Load .env file
	config.LoadEnv()

	// Connect database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	// init logger
	logger := config.NewLogger()

	// Fiber App Init
	app := config.NewFiber(logger)

	// init validator
	validator := config.NewValidator()

	// load routes
	route.LoadRoutes(app, db, logger, validator)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}

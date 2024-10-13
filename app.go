package main

import (
	"amartha-billing-app/common"
	"amartha-billing-app/routes"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(os.Getenv("ENV_LOC"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Define api wrapper
	app := echo.New()
	app.Validator = &common.CustomValidator{Validator: validator.New()}

	// Define the api routes
	routes.DefineApiRoutes(app)

	// Start the http server
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}

package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4/middleware"

	"backend/routers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	log.Print("Starting the service")

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	e := echo.New()
	e.Use(middleware.CORS())

	routers.Router(e)

	log.Print("The service is ready to listen and serve.")
	e.Logger.Fatal(e.Start(":1324"))
}

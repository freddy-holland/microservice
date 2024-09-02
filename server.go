package main

import (
	"log"
	"net/http"
	"os"

	"fholl.net/microservice-base/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	cfg := database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	if err := database.Setup(&cfg); err != nil {
		log.Fatal("Could not connect to database.")
		os.Exit(1)
	}

	// Optionally migrate defined models.
	// database.DB.AutoMigrate()

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/__health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.GET("/__metrics", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

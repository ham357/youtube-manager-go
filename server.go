package main

import (
	"github.com/ham357/youtube-manager-go/middlewares"
	"github.com/ham357/youtube-manager-go/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error loading .env")
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middlewares.YouTubeService())
	e.Use(middlewares.DatabaseService())
	e.Use(middlewares.Firebase())
	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}

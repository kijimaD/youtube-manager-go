package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"youtube-manager/routes"
)

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())

	// Routes
	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

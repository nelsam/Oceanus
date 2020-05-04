package main

import (
	"os"

	"github.com/labstack/echo/v4"
	echoLog "github.com/labstack/gommon/log"
	middleware "github.com/neko-neko/echo-logrus/v2"
	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
)

const (
	url = "http://localhost"
)

func main() {
	// Echo instance
	e := echo.New()

	// Logger
	log.Logger().SetOutput(os.Stdout)
	log.Logger().SetLevel(echoLog.INFO)
	log.Logger().SetFormatter(&logrus.TextFormatter{})
	e.Logger = log.Logger()
	e.Use(middleware.Logger())
	log.Info("Logger enabled!!")

	// Routes
	e.GET("/", hello)
	e.POST("/image", postImageHandler)
	e.GET("/image", getImageHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

package main

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	mw Middlewares
)

func init() {
	mw = Middlewares{}
}

func server() {
	var app *fiber.App = fiber.New()
	app.Use(logger.New())

	app.Get("/", mw.Home)
	app.Post("/ping", mw.Ping)

	if Hostname != "MacBook-Pro.local" {
		app.Use("/file", filesystem.New(filesystem.Config{
			Root:   http.Dir("/tmp"),
			Browse: true,
		}))
	}

	if !IsTestMode { // if this is not in test mode
		app.Listen(os.Getenv("ServerListenPort"))
	} else { // if is test mode
		app.Listen(os.Getenv("ServerListenDevPort"))
	}
}

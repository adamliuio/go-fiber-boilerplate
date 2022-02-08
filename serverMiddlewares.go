package main

import (
	"github.com/gofiber/fiber/v2"
)

type Middlewares struct{}

func (mw Middlewares) Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func (mw Middlewares) Ping(c *fiber.Ctx) error { return c.SendString("pong 👋!") }

package config

import (
	"github.com/gofiber/fiber/v2"
	"synapsis-challange/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}

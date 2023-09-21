package exception

import "github.com/gofiber/fiber/v2"

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

func NewError(statusError int, error string) error {
	return fiber.NewError(statusError, error)
}

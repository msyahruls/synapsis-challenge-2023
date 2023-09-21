package middleware

import (
	"strings"

	"synapsis-challange/helper"
	"synapsis-challange/model/web"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("token")

	if _, _, err := helper.ParseJwt(cookie); err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return c.Status(401).JSON(web.WebResponse{
				Code:    99281,
				Status:  false,
				Message: "token expired",
			})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(web.WebResponse{
			Code:    fiber.StatusUnauthorized,
			Status:  false,
			Message: "unauthorized",
		})
	}

	return c.Next()
}

func IsAdmin(c *fiber.Ctx) bool {
	cookie := c.Cookies("token")

	if _, level, err := helper.ParseJwt(cookie); err != nil || level != "1" {
		return false
	}

	return true
}

func IsCoordinator(c *fiber.Ctx) bool {
	cookie := c.Cookies("token")

	if _, level, err := helper.ParseJwt(cookie); err != nil || level != "2" {
		return false
	}

	return true
}

func IsRelawan(c *fiber.Ctx) bool {
	cookie := c.Cookies("token")

	if _, level, err := helper.ParseJwt(cookie); err != nil || level != "3" {
		return false
	}

	return true
}

func IsSameUserId(c *fiber.Ctx, userId string) bool {
	cookie := c.Cookies("token")

	if user, _, err := helper.ParseJwt(cookie); err != nil || user != userId {
		return false
	}

	return true
}

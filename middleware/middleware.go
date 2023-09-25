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

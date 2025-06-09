package Middleware

import (
	"time"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2"
)

var RateLimit = limiter.New(limiter.Config{
	Max:        200,
	Expiration: 1 * time.Minute,
	KeyGenerator: func(c *fiber.Ctx) string {
		return c.IP()
	},
	LimitReached: func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"error": "Rate limit",
		})
	},
})
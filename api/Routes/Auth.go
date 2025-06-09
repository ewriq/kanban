package Routes

import (
	Handler "auth-api/Handler"
	"github.com/gofiber/fiber/v2"
)

func Auth(app fiber.Router) {
	app.Post("/login", Handler.Login)
	app.Post("/register", Handler.Register)
	app.Post("/user", Handler.User)
}
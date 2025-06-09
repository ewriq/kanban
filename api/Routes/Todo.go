package Routes

import (
	Handler "auth-api/Handler"
	"github.com/gofiber/fiber/v2"
)

func Todo(app fiber.Router) {
	app.Post("/add", Handler.TodoAdd)
	app.Post("/del", Handler.TodoDel)
	app.Post("/list", Handler.TodoList)
}
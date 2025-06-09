package Routes

import (
	Handler "auth-api/Handler"
	"github.com/gofiber/fiber/v2"
)

func Column(app fiber.Router) {
	app.Post("/add", Handler.ColumnAdd)
	app.Post("/del", Handler.ColumnDel)
	app.Post("/list", Handler.ColumnList)
}
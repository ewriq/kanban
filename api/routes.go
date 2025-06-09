package main

import (
	Handler "auth-api/Handler"
	Middleware "auth-api/Middleware"
	"auth-api/Routes"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func Initalize(app *fiber.App){
	app.Use(Middleware.Cors)
	app.Use(Middleware.RateLimit)
	app.Use(helmet.New())
	app.Use(Middleware.Compress)

	auth := app.Group("/api/auth/")
	column := app.Group("/api/column/")
	todo := app.Group("/api/todo/")


	app.Get("/", Handler.Home)

	Routes.Auth(auth)
	Routes.Column(column)
	Routes.Todo(todo)

	app.Use(Middleware.Notfound)
}
package routes

import (
	"github.com/codergab/gojwt/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.GetUser)
	app.Post("/api/logout", controllers.Logout)
}

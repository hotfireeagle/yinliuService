package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	ylApiModule := app.Group("/ylApi")
	serviceModule := ylApiModule.Group("/ope")

	serviceUserModule := serviceModule.Group("/user")
	serviceUserModule.Post("/login", OpsUserDoLogin)
	serviceUserModule.Post("/create_shadow_not_public12131213", OpsUserDoCreate)
}

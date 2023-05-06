package routes

import (
	"github.com/RianIhsan/go-rest-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Get("/api/reads", controllers.ReadAll)
	app.Get("/api/read/:id", controllers.Read)
	app.Post("/api/create", controllers.Create)
	app.Put("/api/update/:id", controllers.Update)
	app.Delete("/api/delete/:id", controllers.Delete)
}

package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhrobby/go-authentication/controllers"
	"github.com/muhrobby/go-authentication/controllers/auth"
	"github.com/muhrobby/go-authentication/middleware"
)

func RouterInit(app *fiber.App) {

	app.Get("/", middleware.Protected, controllers.Home)

	api := app.Group("/api")
	api.Post("/role", controllers.CreateRole)
	api.Get("/roles", controllers.ShowRole)
	api.Delete("/role/:id", controllers.RoleDestroy)
	api.Put("/role/:id", controllers.RoleUpdate)

	Auth := app.Group("/api/auth")
	Auth.Post("/register", auth.Register)
	Auth.Post("/login", auth.Login)

}

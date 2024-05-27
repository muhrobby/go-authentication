package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhrobby/go-authentication/database"
	"github.com/muhrobby/go-authentication/router"
)

func main() {

	app := fiber.New()
	database.ConnectDB()
	router.RouterInit(app)

	app.Listen(":3030")

}

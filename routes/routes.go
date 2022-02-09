package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/inventory_api/controller"
	"github.com/nikitamirzani323/inventory_api/middleware"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())

	app.Post("/api/login", controller.CheckLogin)
	app.Post("/api/loginother", controller.CheckLoginOtherWebsite)

	app.Post("/api/admin", middleware.JWTProtected(), controller.Adminhome)
	// app.Post("/api/admin", controller.Adminhome)

	return app
}

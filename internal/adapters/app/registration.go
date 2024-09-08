package app

import (
	"github.com/billowdev/go-fiber-e-commerce/internal/adapters/http/routers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AppContainer(app *fiber.App, db *gorm.DB) *fiber.App {
	v1 := app.Group("/v1")
	route := routers.NewRoute(v1)
	_ = route
	// SystemFieldApp(route, db)
	return app
}

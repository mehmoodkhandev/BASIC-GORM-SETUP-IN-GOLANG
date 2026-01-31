package router

import (
	"github.com/MehmoodNadeemKhan1/CRUD_API/handler"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(authHandler *handler.AuthHandler) *fiber.App {
	router := fiber.New()

	router.Route("/auth", func(router fiber.Router) {
		router.Post("/", authHandler.Create)
		router.Get("", authHandler.FindAll)
	})

	router.Route("/auth/:id", func(router fiber.Router) {

		router.Put("", authHandler.Update)
		router.Delete("", authHandler.Delete)
		router.Get("", authHandler.FindByID)

	})

	return router

}

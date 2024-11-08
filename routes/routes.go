package routes

import (
    "techs/internal/handler"
    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/users", handler.GetUsers)
//    app.Post("/login", handler.Login)
}

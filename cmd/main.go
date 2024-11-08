package main

import (
    "log"
    "techs/config"
    "techs/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Configuración y carga de variables de entorno
    config.Load()

    // Inicializa Fiber
    app := fiber.New()

    // Configura las rutas
    routes.SetupRoutes(app)

    // Inicia el servidor
    log.Fatal(app.Listen(":3000"))
}

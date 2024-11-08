package handler

import (

    "github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
    users := []string{"Alice", "Bob", "Charlie"}
    //    if err != nil {
      //  return c.Status(500).SendString("Error fetching users")
   // }
    return c.JSON(users)
}

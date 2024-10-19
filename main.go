package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Middleware for access control
func accessControl(c *fiber.Ctx) error {
    // Example: Check for a simple Authorization header
    authHeader := c.Get("Authorization")

    // Dummy authorization logic (replace with real auth)
    if authHeader != "Bearer valid-token" {
        return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized Access")
    }

    // If authorized, continue to the next middleware/handler
    return c.Next()
}

func main() {
    // Create a new Fiber app
    app := fiber.New()

    // Serve static files from the "public" directory
    app.Static("/", "./public")

    // Define route to download a specific file from "public" directory
    app.Get("/download/:filename", func(c *fiber.Ctx) error {
        filename := c.Params("filename")
        err := c.Download("./public/" + filename)

        if err != nil {
            return c.Status(fiber.StatusNotFound).SendString("File not found")
        }
        return nil
    })

    // Apply access control middleware to restricted routes
    app.Use("/restricted", accessControl)

    // Serve files from the restricted directory (only accessible to authorized users)
    app.Static("/restricted", "./restricted-files")

    // Route to download a specific file from the restricted directory
    app.Get("/restricted/download/:filename", func(c *fiber.Ctx) error {
        filename := c.Params("filename")
        err := c.Download("./restricted-files/" + filename)

        if err != nil {
            return c.Status(fiber.StatusNotFound).SendString("File not found")
        }
        return nil
    })

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}

package main

import (
	"log"

	"github.com/alifnuryana/go-auth-jwt/database"
	"github.com/alifnuryana/go-auth-jwt/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)
	database.SetupDatabase()

	log.Fatalln(app.Listen(":4000"))
}

package routes

import (
	"github.com/alifnuryana/go-auth-jwt/controllers"
	"github.com/alifnuryana/go-auth-jwt/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v4"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	api := app.Group("api")

	auth := api.Group("auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)

	todo := api.Group("todo", middleware.Protected())
	todo.Get("/:id", controllers.GetTodo)
	todo.Get("/", controllers.GetTodos)
	todo.Post("/", controllers.CreateTodo)
	todo.Put("/:id", controllers.UpdateTodo)
	todo.Delete("/:id", controllers.DeleteTodo)

	api.Get("/protected", middleware.Protected(), func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		username := claims["Username"].(string)
		return c.SendString("Haloo selamat datang " + username)
	})
}

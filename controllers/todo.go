package controllers

import (
	"fmt"
	"strconv"

	"github.com/alifnuryana/go-auth-jwt/database"
	"github.com/alifnuryana/go-auth-jwt/models"
	"github.com/alifnuryana/go-auth-jwt/repository"
	"github.com/gofiber/fiber/v2"
)

func GetTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInteger, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	todoRepository := repository.NewTodoRepository(database.DB)
	todo, err := todoRepository.GetTodo(idInteger)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "failed",
			"code":    fiber.StatusNotFound,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"code":   fiber.StatusOK,
		"data":   todo,
	})
}

func GetTodos(ctx *fiber.Ctx) error {
	todoRepository := repository.NewTodoRepository(database.DB)
	todos, err := todoRepository.GetTodos()
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "failed",
			"code":    fiber.StatusNotFound,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"code":   fiber.StatusOK,
		"data":   todos,
	})
}

func CreateTodo(ctx *fiber.Ctx) error {
	var newTodo models.Todo

	err := ctx.BodyParser(&newTodo)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	todoRepository := repository.NewTodoRepository(database.DB)
	err = todoRepository.CreateTodo(newTodo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"code":    fiber.StatusCreated,
		"message": "successfully created todo",
	})
}

func UpdateTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	var todo models.Todo

	err = ctx.BodyParser(&todo)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	todoRepository := repository.NewTodoRepository(database.DB)
	err = todoRepository.UpdateTodo(todo, idInt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "failed",
			"code":    fiber.StatusNotFound,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"code":    fiber.StatusOK,
		"message": fmt.Sprintf("successfully updated todo id : %d", idInt),
	})
}

func DeleteTodo(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	todoRepository := repository.NewTodoRepository(database.DB)
	err = todoRepository.DeleteTodo(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "failed",
			"code":    fiber.StatusNotFound,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"code":    fiber.StatusOK,
		"message": "successfully deleted todo",
	})
}

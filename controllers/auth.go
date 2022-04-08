package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/alifnuryana/go-auth-jwt/database"
	"github.com/alifnuryana/go-auth-jwt/helpers"
	"github.com/alifnuryana/go-auth-jwt/models"
	"github.com/alifnuryana/go-auth-jwt/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Register(ctx *fiber.Ctx) error {
	var user models.User
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	userRepository := repository.NewUserRepository(database.DB)
	err = userRepository.CreateUser(user)
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
		"message": "successfully created user",
	})
}

func Login(ctx *fiber.Ctx) error {
	var requestLogin models.RequestLogin

	err := ctx.BodyParser(&requestLogin)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	userRepository := repository.NewUserRepository(database.DB)
	var user models.User
	if strings.Contains(requestLogin.Identity, "@") {
		user, err = userRepository.GetUserByEmail(requestLogin.Identity)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "failed",
				"code":    fiber.StatusInternalServerError,
				"message": "your username / email is not valid",
			})
		}
	} else {
		user, err = userRepository.GetUserByUsername(requestLogin.Identity)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "failed",
				"code":    fiber.StatusInternalServerError,
				"message": "your username / email is not valid",
			})
		}
	}

	if requestLogin.Password != user.Password {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "failed",
			"code":    fiber.StatusUnauthorized,
			"message": "your password is not valid",
		})
	}

	stringId := strconv.Itoa(int(user.ID))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.CustomClaim{
		Username: user.Username,
		Role:     user.Role,
		Id:       stringId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-auth-jwt",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
		},
	})

	tokenString, err := token.SignedString([]byte(helpers.Load("JWT_KEY")))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"code":   fiber.StatusOK,
		"token":  tokenString,
	})
}

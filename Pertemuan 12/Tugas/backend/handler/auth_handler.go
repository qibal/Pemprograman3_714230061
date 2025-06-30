package handler

import (
	"fmt"
	"inibackend/config/middleware"
	"inibackend/model"
	pwd "inibackend/pkg/password"
	"inibackend/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Login godoc
// @Summary Login User
// @Description Melakukan proses login dan mengembalikan token PASETO jika username dan password valid.
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body model.LoginRequest true "Login credentials (username dan password)"
// @Success 200 {object} model.LoginResponse "Login success"
// @Failure 400 "Invalid body"
// @Failure 401 "Username not found or Wrong password"
// @Failure 500 "Failed to generate token"
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	var req model.UserLogin

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
	}

	user, err := repository.FindUserByUsername(c.Context(), req.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Username not found"})
	}

	// Cek password input hash yang tersimpan
	if !pwd.CheckPasswordHash(req.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong password"})
	}

	// Generate token PASETO
	token, err := middleware.EncodeWithRoleHours(user.Role, user.Username, 2)
	if err != nil {
		fmt.Println("Token generation error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{
		"message": "Login success",
		"token":   token,
	})
}

// Register godoc
// @Summary Register a new user
// @Description Mendaftarkan pengguna baru dengan username, password, dan role.
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body model.UserLogin true "User Registration Details"
// @Success 201 {object} model.RegisterResponse "User registered successfully"
// @Failure 400 {object} object "Invalid input or missing fields"
// @Failure 409 {object} object "Username already exists"
// @Failure 500 {object} object "Failed to hash password"
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	var req model.UserLogin

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if req.Username == "" || req.Password == "" || req.Role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username, password, and role are required"})
	}

	hashed, err := pwd.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	req.Password = hashed

	id, err := repository.InsertUser(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}

	var insertedID string
	if oid, ok := id.(primitive.ObjectID); ok {
		insertedID = oid.Hex()
	}

	response := model.RegisterResponse{
		Message: "User registered successfully",
		ID:      insertedID,
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

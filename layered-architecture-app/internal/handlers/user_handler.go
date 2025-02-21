package handlers

import (
	"internal/models"
	"internal/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) handleGetUsers(c *fiber.Ctx) error {
	user := h.userService.GetUsers()

	return c.JSON(user)
}

func (h *UserHandler) handleGetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("can not parse id")
	}

	user, err := h.userService.GetUser(id)
	if err != nil {
		return c.Status(404).SendString("user not found")
	}

	return c.JSON(user)
}

func (h *UserHandler) handleCreateUser(c *fiber.Ctx) error {
	userRequest := models.UserCreateRequest{}

	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(400).SendString("Invalid user data")
	}

	err := userRequest.Validate()
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	user := h.userService.CreateUser(userRequest)

	return c.JSON(user)
}

func (h *UserHandler) handleUpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("can not parse id")
	}

	updateRequest := models.UserUpdateRequest{}
	if err := c.BodyParser(&updateRequest); err != nil {
		return c.Status(400).SendString("Invalid user data")
	}

	user, err := h.userService.UpdateUser(id, updateRequest)
	if err != nil {
		return c.Status(404).SendString("user not found")
	}

	return c.JSON(user)
}

func (h *UserHandler) handleDeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("can not parse id")
	}

	h.userService.DeleteUser(id)

	return c.SendString("User deleted")
}

func (h *UserHandler) SetRoutes(app *fiber.App) {
	app.Get("/users", h.handleGetUsers)
	app.Get("/users/:id", h.handleGetUser)
	app.Post("/users", h.handleCreateUser)
	app.Put("/users/:id", h.handleUpdateUser)
	app.Delete("/users/:id", h.handleDeleteUser)
}

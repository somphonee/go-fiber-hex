package handlers

import (
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/somphonee/go-fiber-hex/internal/core/domain"
	"github.com/somphonee/go-fiber-hex/internal/core/ports"
	"github.com/somphonee/go-fiber-hex/pkg/errors"

)


type UserHandler struct {
	userService ports.UserService
	validator   *validator.Validate
}

func NewUserHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator.New(),
	}
}
func (h *UserHandler) Create(c *fiber.Ctx) error {
	req := new(domain.CreateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.NewError("invalid request", err))
	}

	if err := h.validator.Struct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errors.NewValidationError(err))
	}

	user, err := h.userService.Create(c.Context(), req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(errors.NewError("failed to create user", err))
	}

	return c.Status(http.StatusCreated).JSON(user)
}
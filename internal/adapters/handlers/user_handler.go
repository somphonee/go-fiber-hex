package handlers

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/somphonee/go-fiber-hex/internal/core/domain"
	"github.com/somphonee/go-fiber-hex/internal/core/ports"
	"github.com/somphonee/go-fiber-hex/pkg/response"
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

// Add this function to your handler struct
func (h *UserHandler) translateErrors(err error) map[string]string {
	result := make(map[string]string)
	
	if errors, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errors {
					field := err.Field()
					switch err.Tag() {
					case "required":
							result[field] = field + " is required"
					case "email":
							result[field] = field + " should be a valid email"
					// Add more cases as needed
					default:
							result[field] = field + " is invalid"
					}
			}
	}
	
	return result
}
func (h *UserHandler) Create(c *fiber.Ctx) error {
	req := new(domain.CreateUserRequest)
	if err := c.BodyParser(req); err != nil {
			return response.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	// Validate request
	if err := h.validator.Struct(req); err != nil {
    var validationErrors []string
    for _, err := range err.(validator.ValidationErrors) {
        validationErrors = append(validationErrors, err.Field() + ": " + err.Tag())
    }
    return response.SendError(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
}

  err := h.userService.Create(c.Context(), req)
	if err != nil {
			return response.SendError(c, fiber.StatusInternalServerError, "Failed to create user", err.Error())
	}

	return response.SendCreated(c, "ok", "User created successfully")
}
func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10 , 32)
	if  err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "Invalid ID", err.Error())
	}
	user, err := h.userService.GetByID(c.Context(), uint(id))
	if err != nil {
		return response.SendError(c, fiber.StatusNotFound, "User not found", err.Error())
	}
	return response.SendData(c, user,"User retrieved successfully") 		
}

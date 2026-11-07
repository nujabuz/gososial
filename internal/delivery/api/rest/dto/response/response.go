package response

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/nujabuz/gososial/internal/domain/entities"
)

type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func ResponseSuccess(ctx fiber.Ctx, status int, message string, data any) error {
	return ctx.Status(status).JSON(response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ResponseFailed(ctx fiber.Ctx, status int, message string, error any) error {
	return ctx.Status(status).JSON(response{
		Success: false,
		Message: message,
		Error:   error,
	})
}

func ResponseServiceError(ctx fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, entities.ErrValidation):
		return ResponseFailed(ctx, fiber.StatusBadRequest, err.Error(), fiber.ErrBadRequest)
	case errors.Is(err, entities.ErrNotFound):
		return ResponseFailed(ctx, fiber.StatusNotFound, err.Error(), fiber.ErrNotFound)
	case errors.Is(err, entities.ErrAlreadyExists):
		return ResponseFailed(ctx, fiber.StatusConflict, err.Error(), fiber.ErrConflict)
	default:
		return ResponseFailed(
			ctx,
			fiber.StatusInternalServerError,
			entities.ErrInternalServer.Error(),
			fiber.ErrInternalServerError,
		)
	}
}

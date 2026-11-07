// Package controllers
package controllers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nujabuz/gososial/internal/application/interfaces"
	"github.com/nujabuz/gososial/internal/delivery/api/rest/dto/request"
	"github.com/nujabuz/gososial/internal/delivery/api/rest/dto/response"
)

type UserController struct {
	service interfaces.UserService
}

func NewUserController(api fiber.Router, userService interfaces.UserService) *UserController {
	controller := &UserController{
		service: userService,
	}

	auth := api.Group("auth")
	auth.Post("/register", controller.Register)

	return controller
}

func (uc *UserController) Register(ctx fiber.Ctx) error {
	requestRegisterUser := new(request.RequestRegisterUser)
	if err := ctx.Bind().Body(requestRegisterUser); err != nil {
		return response.ResponseFailed(ctx, fiber.StatusBadRequest, "bad request for register", err)
	}

	result, err := uc.service.CreateUser(ctx, requestRegisterUser.ToParam())
	if err != nil {
		return response.ResponseServiceError(ctx, err)
	}

	return response.ResponseSuccess(
		ctx, fiber.StatusCreated,
		"successfully registered user",
		response.NewResponseUser(result),
	)
}

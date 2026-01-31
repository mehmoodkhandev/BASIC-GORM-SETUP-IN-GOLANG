package handler

import (
	"strconv"

	"github.com/MehmoodNadeemKhan1/CRUD_API/data/request"
	"github.com/MehmoodNadeemKhan1/CRUD_API/data/response"
	"github.com/MehmoodNadeemKhan1/CRUD_API/helper"
	"github.com/MehmoodNadeemKhan1/CRUD_API/services"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authServices services.AuthService
}

func NewAuthHandler(Services services.AuthService) *AuthHandler {
	return &AuthHandler{
		authServices: Services,
	}

}

func (handler *AuthHandler) Create(ctx *fiber.Ctx) error {

	createAuthRequest := request.AuthCreateRequest{}
	err := ctx.BodyParser(&createAuthRequest)
	helper.ErrorPanic(err)
	handler.authServices.Create(createAuthRequest)
	webResponse := response.Response{
		Code:    201,
		Status:  "success",
		Message: "User created successfully",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (handler *AuthHandler) Update(ctx *fiber.Ctx) error {

	updateAuthRequest := request.AuthUpdateRequest{}
	err := ctx.BodyParser(&updateAuthRequest)
	helper.ErrorPanic(err)
	authId := ctx.Params("id")
	id, err := strconv.Atoi(authId)
	helper.ErrorPanic(err)
	updateAuthRequest.ID = uint(id)
	handler.authServices.Update(updateAuthRequest)
	webResponse := response.Response{
		Code:    200,
		Status:  "success",
		Message: "User updated successfully",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (handler *AuthHandler) Delete(ctx *fiber.Ctx) error {
	authId := ctx.Params("id")
	id, err := strconv.Atoi(authId)
	helper.ErrorPanic(err)
	handler.authServices.Delete(id)
	webResponse := response.Response{
		Code:    200,
		Status:  "success",
		Message: "User deleted successfully",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (handler *AuthHandler) FindByID(ctx *fiber.Ctx) error {
	noteId := ctx.Params("id")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)
	authResponse := handler.authServices.FindById(id)
	webResponse := response.Response{
		Code:    200,
		Status:  "success",
		Message: "User retrieved successfully",
		Data:    authResponse,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (handler *AuthHandler) FindAll(ctx *fiber.Ctx) error {
	authResponse := handler.authServices.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "success",
		Message: "Users retrieved successfully",
		Data:    authResponse,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}

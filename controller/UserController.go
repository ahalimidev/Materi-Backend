package controller

import (
	"fmt"
	"materi/helper"
	"materi/model/request"
	"materi/service"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController interface {
	All(context *fiber.Ctx) error
	FindByID(context *fiber.Ctx) error
	Create(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
	Delete(context *fiber.Ctx) error
	Upload(context *fiber.Ctx) error
}

type userController struct {
	userService service.UserService
}

// All implements UserController
func (c *userController) All(context *fiber.Ctx) error {
	user, err := c.userService.All(context.Context())

	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to get data", http.StatusOK, "success", user)
	return context.Status(http.StatusOK).JSON(response)
}

// Create implements UserController
func (c *userController) Create(context *fiber.Ctx) error {
	var inputData request.UserCreate
	if err := context.BodyParser(&inputData); err != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	errors := request.ValidateStruct(&inputData)
	if errors != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", errors)
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	_, err := c.userService.Create(context.Context(), inputData)
	if err != nil {
		response := helper.APIResponse("Failed to post data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to post data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// Delete implements UserController
func (c *userController) Delete(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_user"))
	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	user, err := c.userService.FindByID(context.Context(), strconv.Itoa(id))

	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	if user.Foto.String != "" {
		path := fmt.Sprintf("file/user/%s", user.Foto.String)
		_, err := os.Stat(path)
		if !os.IsNotExist(err) {
			os.Remove(path) // remove a single file
		}
	}
	c.userService.Delete(context.Context(), strconv.Itoa(id))
	response := helper.APIResponse("Success to delete data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// FindByID implements UserController
func (c *userController) FindByID(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_user"))
	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	user, err := c.userService.FindByID(context.Context(), strconv.Itoa(id))

	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse("Success to get data", http.StatusOK, "success", user)
	return context.Status(http.StatusOK).JSON(response)
}

// Update implements UserController
func (c *userController) Update(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_user"))
	if err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	var inputData request.UserUpdate

	if err := context.BodyParser(&inputData); err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	inputData.IdUser = strconv.Itoa(id)

	errors := request.ValidateStruct(&inputData)
	if errors != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", errors)
		return context.Status(http.StatusBadRequest).JSON(response)
	}

	if err := c.userService.Update(context.Context(), inputData); err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success to put data", http.StatusOK, "success", nil)
	return context.Status(http.StatusOK).JSON(response)
}

// Upload implements UserController
func (c *userController) Upload(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id_user"))
	if err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	var inputData request.UserProfile

	if err := context.BodyParser(&inputData); err != nil {
		response := helper.APIResponse("Failed to put data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	inputData.IdUser = strconv.Itoa(id)

	result, err := c.userService.FindByID(context.Context(), strconv.Itoa(id))

	if err != nil {
		response := helper.APIResponse("Failed to get data", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	if result.Foto.String != "" {
		path := fmt.Sprintf("file/user/%s", result.Foto.String)
		_, err := os.Stat(path)
		if !os.IsNotExist(err) {
			os.Remove(path) // remove a single file
		}
	}
	file, err := context.FormFile("foto")
	if err != nil || file.Size == 0 {
		inputData.Foto = ""
		response := helper.APIResponse("Failed to put data1", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)
	} else {
		extension := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + extension
		inputData.Foto = newFileName
		path := fmt.Sprintf("file/foto/%s", newFileName)
		err = context.SaveFile(file, path)
		if err != nil {
			response := helper.APIResponse("Failed to upload2", http.StatusBadRequest, "error", err.Error())
			return context.Status(http.StatusBadRequest).JSON(response)

		}
		c.userService.Upload(context.Context(), inputData)
		response := helper.APIResponse("Success to put data", http.StatusOK, "success", nil)
		return context.Status(http.StatusOK).JSON(response)
	}
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

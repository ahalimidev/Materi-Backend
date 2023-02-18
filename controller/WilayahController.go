package controller

import (
	"materi/helper"
	"materi/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type WilayahController interface {
	Provinsi(context *fiber.Ctx) error
	Kabupaten(context *fiber.Ctx) error
	Kecamatan(context *fiber.Ctx) error
	Desa(context *fiber.Ctx) error
}

type wilayahController struct {
	wilayahService service.WilayahService
}

// Desa implements WilayahController
func (c *wilayahController) Desa(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("kecamatan_id"))
	if err != nil {
		response := helper.APIResponse("Failed to get", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	result, err := c.wilayahService.Desa(context.Context(), strconv.Itoa(id))
	if err != nil {
		response := helper.APIResponse("Failed to get", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	response := helper.APIResponse("wilayah", http.StatusOK, "success", result)
	return context.Status(http.StatusOK).JSON(response)
}

// Kabupaten implements WilayahController
func (c *wilayahController) Kabupaten(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("provinsi_id"))
	if err != nil {
		response := helper.APIResponse("Failed to get", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	result, err := c.wilayahService.Kabupaten(context.Context(), strconv.Itoa(id))
	if err != nil {
		response := helper.APIResponse("Failed to get", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	response := helper.APIResponse("wilayah", http.StatusOK, "success", result)
	return context.Status(http.StatusOK).JSON(response)
}

// Kecamatan implements WilayahController
func (c *wilayahController) Kecamatan(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("kabupaten_id"))
	if err != nil {
		response := helper.APIResponse("Failed to get", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	result, err := c.wilayahService.Kecamatan(context.Context(), strconv.Itoa(id))
	if err != nil {
		response := helper.APIResponse("Failed to get", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	response := helper.APIResponse("wilayah", http.StatusOK, "success", result)
	return context.Status(http.StatusOK).JSON(response)
}

// Provinsi implements WilayahController
func (c *wilayahController) Provinsi(context *fiber.Ctx) error {
	result, err := c.wilayahService.Provinsi(context.Context())
	if err != nil {
		response := helper.APIResponse("Failed to get", http.StatusBadRequest, "error", err.Error())
		return context.Status(http.StatusBadRequest).JSON(response)

	}
	response := helper.APIResponse("wilayah", http.StatusOK, "success", result)
	return context.Status(http.StatusOK).JSON(response)
}

func NewWilayahController(wilayahService service.WilayahService) WilayahController {
	return &wilayahController{
		wilayahService: wilayahService,
	}
}

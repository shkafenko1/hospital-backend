package handlers

import (
	"hospital-backend/internal/domain/models"
	"hospital-backend/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type WardHandler struct {
	service *service.WardService
}

func NewWardHandler(service *service.WardService) *WardHandler {
	return &WardHandler{service: service}
}

func (h *WardHandler) Create(c fiber.Ctx) error {
	var ward models.Ward
	if err := c.Bind().Body(&ward); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	
	if err := h.service.Create(&ward); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(201).JSON(ward)
}

func (h *WardHandler) GetByID(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	ward, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Ward not found"})
	}
	
	return c.JSON(ward)
}

func (h *WardHandler) GetAll(c fiber.Ctx) error {
	wards, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.JSON(wards)
}

func (h *WardHandler) Update(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	var ward models.Ward
	if err := c.Bind().Body(&ward); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	
	ward.ID = id
	if err := h.service.Update(&ward); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.JSON(ward)
}

func (h *WardHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	if err := h.service.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(204).Send(nil)
}

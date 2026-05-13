package handlers

import (
	"hospital-backend/internal/domain/models"
	"hospital-backend/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type HospitalizationHandler struct {
	service *service.HospitalizationService
}

func NewHospitalizationHandler(service *service.HospitalizationService) *HospitalizationHandler {
	return &HospitalizationHandler{service: service}
}

func (h *HospitalizationHandler) Create(c fiber.Ctx) error {
	var hospitalization models.Hospitalization
	if err := c.Bind().Body(&hospitalization); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	
	if err := h.service.Create(&hospitalization); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(201).JSON(hospitalization)
}

func (h *HospitalizationHandler) GetByID(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	hospitalization, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Hospitalization not found"})
	}
	
	return c.JSON(hospitalization)
}

func (h *HospitalizationHandler) GetAll(c fiber.Ctx) error {
	hospitalizations, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.JSON(hospitalizations)
}

func (h *HospitalizationHandler) Update(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	var hospitalization models.Hospitalization
	if err := c.Bind().Body(&hospitalization); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	
	hospitalization.ID = id
	if err := h.service.Update(&hospitalization); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.JSON(hospitalization)
}

func (h *HospitalizationHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	if err := h.service.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(204).Send(nil)
}

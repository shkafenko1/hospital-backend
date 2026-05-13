package handlers

import (
	"hospital-backend/internal/domain/models"
	"hospital-backend/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type VisitHandler struct {
	service *service.VisitService
}

func NewVisitHandler(service *service.VisitService) *VisitHandler {
	return &VisitHandler{service: service}
}

func (h *VisitHandler) Create(c fiber.Ctx) error {
	var visit models.Visit
	if err := c.Bind().Body(&visit); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	
	if err := h.service.Create(&visit); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(201).JSON(visit)
}

func (h *VisitHandler) GetByID(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	visit, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Visit not found"})
	}
	
	return c.JSON(visit)
}

func (h *VisitHandler) GetAll(c fiber.Ctx) error {
	visits, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.JSON(visits)
}

func (h *VisitHandler) Update(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	var visit models.Visit
	if err := c.Bind().Body(&visit); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	
	visit.ID = id
	if err := h.service.Update(&visit); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.JSON(visit)
}

func (h *VisitHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	if err := h.service.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(204).Send(nil)
}

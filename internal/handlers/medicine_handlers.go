package handlers

import (
	"hospital-backend/internal/domain/models"
	"hospital-backend/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type MedicineHandler struct {
	service *service.MedicineService
}

func NewMedicineHandler(service *service.MedicineService) *MedicineHandler {
	return &MedicineHandler{service: service}
}

func (h *MedicineHandler) Create(c fiber.Ctx) error {
	var medicine models.Medicine
	if err := c.Bind().Body(&medicine); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	
	if err := h.service.Create(&medicine); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(201).JSON(medicine)
}

func (h *MedicineHandler) GetByID(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	medicine, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Medicine not found"})
	}
	
	return c.JSON(medicine)
}

func (h *MedicineHandler) GetAll(c fiber.Ctx) error {
	medicines, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.JSON(medicines)
}

func (h *MedicineHandler) Update(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	var medicine models.Medicine
	if err := c.Bind().Body(&medicine); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	
	medicine.ID = id
	if err := h.service.Update(&medicine); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.JSON(medicine)
}

func (h *MedicineHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	if err := h.service.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(204).Send(nil)
}

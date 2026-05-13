package handlers

import (
	"hospital-backend/internal/domain/models"
	"hospital-backend/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type DoctorHandler struct {
	service *service.DoctorService
}

func NewDoctorHandler(service *service.DoctorService) *DoctorHandler {
	return &DoctorHandler{service: service}
}

func (h *DoctorHandler) Create(c fiber.Ctx) error {
	var doctor models.Doctor
	if err := c.Bind().Body(&doctor); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	
	if err := h.service.Create(&doctor); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(201).JSON(doctor)
}

func (h *DoctorHandler) GetByID(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	doctor, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Doctor not found"})
	}
	
	return c.JSON(doctor)
}

func (h *DoctorHandler) GetAll(c fiber.Ctx) error {
	doctors, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.JSON(doctors)
}

func (h *DoctorHandler) Update(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	var doctor models.Doctor
	if err := c.Bind().Body(&doctor); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	
	doctor.ID = id
	if err := h.service.Update(&doctor); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.JSON(doctor)
}

func (h *DoctorHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	
	if err := h.service.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(204).Send(nil)
}

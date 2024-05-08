package handler

import (
	"log"

	"github.com/chilah/go-expense-api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type ExpenseHandler struct {
	es port.ExpenseService
}

func NewExpenseHanlder(es port.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		es,
	}
}

func (eh *ExpenseHandler) GetAll(c *fiber.Ctx) error {
	expenses, err := eh.es.GetAll()

	if err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   expenses,
	})
}

func (eh *ExpenseHandler) Create(c *fiber.Ctx) error {
	err := eh.es.Create()

	if err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func (eh *ExpenseHandler) UpdateByID(c *fiber.Ctx) error {
	err := eh.es.UpdateByID()

	if err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Update expense successfully",
	})
}

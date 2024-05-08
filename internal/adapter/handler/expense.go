package handler

import (
	"log"

	"github.com/chilah/go-expense-api/internal/core/domain"
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

type UpdateExpense struct {
	ID     uint `json:"id"`
	Amount int  `json:"amount"`
}

func (eh *ExpenseHandler) UpdateByID(c *fiber.Ctx) error {
	inputExp := UpdateExpense{}

	if err := c.BodyParser(&inputExp); err != nil {
		return err
	}

	expense := domain.Expense{
		ID:     inputExp.ID,
		Amount: inputExp.Amount,
	}

	err := eh.es.UpdateByID(&expense)

	if err != nil {
		// log.Fatal(err)
		// return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		// 	"status": fmt.Sprintf("Unable to find id: %d", inputExp.ID),
		// })

		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Update expense successfully",
	})
}

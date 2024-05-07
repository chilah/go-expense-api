package main

import (
	"fmt"
	"time"

	"github.com/chilah/go-expense-api/internal/adapter/db"
	"github.com/gofiber/fiber/v2"
)

const PORT = ":8080"

type Expense struct {
	ID          uint
	CreateAt    time.Time
	Description string
	Amount      int
	Category    int
	Test        int
}

func main() {
	supabase := db.New()
	db := supabase.Connect()

	defer db.Close()

	app := fiber.New()
	api := app.Group("/api")

	api.Get("/ping", func(c *fiber.Ctx) error {
		expense := &Expense{}

		res, _ := db.Query("SELECT * from expense")

		res.Scan(expense)

		return c.Status(fiber.StatusOK).JSON(expense)
	})

	fmt.Println("hello world")

	app.Listen(PORT)
}

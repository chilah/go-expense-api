package main

import (
	"fmt"

	"github.com/chilah/go-expense-api/internal/adapter/db"
	"github.com/chilah/go-expense-api/internal/adapter/handler"
	"github.com/chilah/go-expense-api/internal/adapter/repository"
	"github.com/chilah/go-expense-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

const PORT = ":8080"

func main() {
	supabase := db.New()
	db := supabase.Connect()

	defer db.Close()

	app := fiber.New()
	api := app.Group("/api")

	expense := api.Group("/expense")
	expenseRepository := repository.NewUserRepository(db)
	expenseService := service.NewExpenseService(expenseRepository)
	expenseHandler := handler.NewExpenseHanlder(expenseService)

	expense.Get("", expenseHandler.GetAll)

	fmt.Printf("Start server on port%s successfully!", PORT)

	app.Listen(PORT)
}

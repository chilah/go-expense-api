package port

import "github.com/chilah/go-expense-api/internal/core/domain"

type ExpenseRepository interface {
	GetAll() (*[]domain.Expense, error)
	Create() error
	UpdateByID() error
}

type ExpenseService interface {
	GetAll() (*[]domain.Expense, error)
	Create() error
	UpdateByID() error
}

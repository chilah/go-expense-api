package service

import (
	"log"

	"github.com/chilah/go-expense-api/internal/core/domain"
	"github.com/chilah/go-expense-api/internal/core/port"
)

type ExpenseService struct {
	er port.ExpenseRepository
}

func NewExpenseService(ur port.ExpenseRepository) *ExpenseService {
	return &ExpenseService{
		ur,
	}
}

func (es *ExpenseService) GetAll() (*[]domain.Expense, error) {
	expenses, err := es.er.GetAll()

	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	return expenses, nil
}

func (es *ExpenseService) Create() error {
	err := es.er.Create()

	return err
}

func (es *ExpenseService) UpdateByID(e *domain.Expense) error {
	_, err := es.er.FindByID(int(e.ID))

	if err != nil {
		return err
	}

	updateErr := es.er.UpdateByID(e)

	if updateErr != nil {
		return err
	}

	return nil
}

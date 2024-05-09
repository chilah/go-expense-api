package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/chilah/go-expense-api/internal/core/domain"
)

type ExpenseRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{
		db,
	}
}

func (ur *ExpenseRepository) GetAll() (*[]domain.Expense, error) {
	rows, err := ur.db.Query("select * from expense")

	if err != nil {
		log.Fatal(err)
	}

	expenses := []domain.Expense{}

	for rows.Next() {
		expense := domain.Expense{}

		if err := rows.Scan(
			&expense.ID,
			&expense.CreateAt,
			&expense.Description,
			&expense.Amount,
			&expense.Category,
			&expense.Test,
		); err != nil {
			log.Fatal(err)
		}

		expenses = append(expenses, expense)
	}

	return &expenses, nil
}

func (er *ExpenseRepository) Create() error {
	_, err := er.db.Exec(
		"insert into expense (description, amount, category_id, test_id) values ($1, $2, $3, $4)",
		"go test", 200, 1, 1,
	)

	if err != nil {
		log.Fatal(err)

		return err
	}

	return nil
}

func (er *ExpenseRepository) UpdateByID(e *domain.Expense) error {

	_, err := er.db.Exec(
		"UPDATE expense SET amount = $1 WHERE id = $2",
		e.Amount, e.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (er *ExpenseRepository) FindByID(id int) (*domain.Expense, error) {
	exp := domain.Expense{}

	err := er.db.QueryRow("SELECT * FROM expense WHERE id = $1;", id).Scan(
		&exp.ID,
		&exp.CreateAt,
		&exp.Description,
		&exp.Amount,
		&exp.Category,
		&exp.Test,
	)

	if err != nil {
		return nil, fmt.Errorf("unalble to find the expense by %d", id)
	}

	return &exp, nil
}

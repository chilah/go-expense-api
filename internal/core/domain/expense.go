package domain

import "time"

type Expense struct {
	ID          uint
	CreateAt    time.Time
	Description string
	Amount      int
	Category    int
	Test        int
}

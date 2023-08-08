package entities

import "time"

type ExpenseRequest struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Date        string  `json:"date"`
	Category    string  `json:"category"`
	Description string  `json:"description,omitempty"`
}

type Expense struct {
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Date        time.Time `json:"date"`
	Category    string    `json:"category"`
	Description string    `json:"description,omitempty"`
}

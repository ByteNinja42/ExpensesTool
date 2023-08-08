package entities

import "time"

type UserSignUp struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

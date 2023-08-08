package service

import (
	"fmt"
	"time"

	"github.com/ByteNinja42/ExpensesTool/internal/entities"
)

type ExpensesService struct {
	ExpensesRepo
}

func NewExpensesService() *ExpensesService {
	return &ExpensesService{}
}

type ExpensesRepo interface {
	AddExpense(entities.Expense, string) error
}

func (ex ExpensesService) CreateExpense(expense entities.ExpenseRequest, userID string) error {
	if userID == "" {
		return entities.ErrUIDEmpty
	}
	timeExpense, err := time.Parse(time.DateTime, expense.Date)
	if err != nil {
		return err
	}
	exp := entities.Expense{
		Name:        expense.Name,
		Price:       expense.Price,
		Date:        timeExpense,
		Category:    expense.Category,
		Description: expense.Description,
	}
	isValid, err := isExpenseValid(exp)
	if !isValid || err != nil {
		return err
	}
	fmt.Println(exp)
	err = ex.AddExpense(exp, userID)
	if err != nil {
		return err
	}
	return nil
}

func isExpenseValid(expense entities.Expense) (bool, error) {
	if expense.Name == "" {
		return false, entities.ErrNameEmpty
	}
	if expense.Price <= 0 {
		return false, entities.ErrPriceDEmpty
	}
	if expense.Category == "" {
		return false, entities.ErrCategoryEmpty
	}
	if expense.Date.IsZero() {
		return false, entities.ErrDateEmpty
	}
	return true, nil
}

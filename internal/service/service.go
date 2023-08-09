package service

import (
	"fmt"
	"regexp"
	"time"

	"github.com/ByteNinja42/ExpensesTool/internal/entities"
	"golang.org/x/crypto/bcrypt"
)

type ExpensesService struct {
	Expense ExpensesRepo
	User    UserRepo
}

func NewExpensesService() *ExpensesService {
	return &ExpensesService{}
}

type ExpensesRepo interface {
	AddExpense(entities.Expense, string) error
}

type UserRepo interface {
	CreateUser(entities.UserSignUp) error
	IsUserExists(email string, passwordHash string) (bool, error)
}

func (ex ExpensesService) UserSignUp(signUp entities.UserSignUp) error {
	isValid, err := isSignUpValid(signUp)
	if !isValid || err != nil {
		return fmt.Errorf("err in service : %w", err)
	}
	passwordHash, err := GenerateHash(signUp.Password)
	if err != nil {
		return fmt.Errorf("err in service : %w", err)
	}
	signUp.Password = passwordHash
	isExist, err := ex.User.IsUserExists(signUp.Email, signUp.Password)
	if isExist {
		return fmt.Errorf("err in service : %w", err)
	}
	err = ex.User.CreateUser(signUp)
	if err != nil {
		return fmt.Errorf("err in service : %w", err)
	}
	return nil
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
	err = ex.Expense.AddExpense(exp, userID)
	if err != nil {
		return err
	}
	return nil
}

func isSignUpValid(signUp entities.UserSignUp) (bool, error) {
	if signUp.FirstName == "" {
		return false, entities.ErrFirstNameEmpty
	}
	if signUp.LastName == "" {
		return false, entities.ErrLastNameEmpty
	}
	if signUp.Password == "" {
		return false, entities.ErrPasswordEmpty
	}
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	if isEmailValid := regexp.MustCompile(emailRegex).MatchString(signUp.Email); !isEmailValid {
		return false, entities.ErrEmailInvalid
	}
	return true, nil
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

func GenerateHash(password string) (string, error) {
	saltedBytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashPassword := string(hashedBytes)
	return hashPassword, nil
}
func CheckPassword(password, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err
}

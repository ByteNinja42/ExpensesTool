package service

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/ByteNinja42/ExpensesTool/internal/config"
	"github.com/ByteNinja42/ExpensesTool/internal/entities"
	"github.com/dgrijalva/jwt-go"
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
	CreateUser(signUp entities.UserSignUp) error
	IsUserExists(email string, passwordHash string) (bool, error)
	GetHashPasswordByEmail(email string) (string, error)
	GetUserIDByEmail(email string) (string, error)
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
	//isExist, err := ex.User.IsUserExists(signUp.Email, signUp.Password)
	//if isExist {
	//	return fmt.Errorf("err in service : %w", err)
	//}
	fmt.Println(signUp)
	//err = ex.User.CreateUser(signUp)
	//if err != nil {
	//	return fmt.Errorf("err in service : %w", err)
	//}
	return nil
}

func (ex ExpensesService) UserSignIn(signIn entities.UserSignIn) (string, error) {
	isValid, err := isSignInvalid(signIn)
	if !isValid || err != nil {
		return "", fmt.Errorf("err in service : %w", err)
	}
	passwordHash, err := ex.User.GetHashPasswordByEmail(signIn.Email)
	if err != nil {
		return "", fmt.Errorf("err in database : %w", err)
	}
	err = isPasswordCorrect(signIn.Password, passwordHash)
	if err != nil {
		return "", fmt.Errorf("err in service : password isn't correct")
	}

	uid, err := ex.User.GetUserIDByEmail(signIn.Email)
	if err != nil {
		return "", err
	}
	token, err := CreateToken(uid)
	if err != nil {
		return "", err
	}
	return token, nil
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

func isSignInvalid(signIn entities.UserSignIn) (bool, error) {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	if isEmailValid := regexp.MustCompile(emailRegex).MatchString(signIn.Email); !isEmailValid {
		return false, entities.ErrEmailInvalid
	}
	if signIn.Password == "" {
		return false, entities.ErrPasswordEmpty
	}
	return true, nil
}

func isPasswordCorrect(password, hashPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err
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

type JWTClaim struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

func CreateToken(userID string) (string, error) {
	timeExp, err := strconv.Atoi(config.GetEnv("TOKEN_EXPIRES", "15"))
	if err != nil {
		return "", err
	}
	claims := &JWTClaim{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(timeExp) * time.Minute).Unix(),
		},
	}
	fmt.Println(claims.ExpiresAt)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

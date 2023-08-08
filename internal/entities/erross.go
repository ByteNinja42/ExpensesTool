package entities

import "errors"

var ErrUIDEmpty = errors.New("user ID can't be empty")
var ErrNameEmpty = errors.New("expense name can't be empty")
var ErrPriceDEmpty = errors.New("expense price can't be empty")
var ErrDateEmpty = errors.New("expense date can't be empty")
var ErrCategoryEmpty = errors.New("expense category can't be empty")

var ErrFirstNameEmpty = errors.New("user first name can't be empty")
var ErrLastNameEmpty = errors.New("user last name can't be empty")
var ErrPasswordEmpty = errors.New("user password can't be empty")
var ErrEmailInvalid = errors.New("email is invalid format")

var ErrUserExists = errors.New("user with this data already exists")

package entities

import "errors"

var ErrUIDEmpty = errors.New("err : user ID can't be empty")
var ErrNameEmpty = errors.New("err : expense name can't be empty")
var ErrPriceDEmpty = errors.New("err : expense price can't be empty")
var ErrDateEmpty = errors.New("err : expense date can't be empty")
var ErrCategoryEmpty = errors.New("err : expense category can't be empty")

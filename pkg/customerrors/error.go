package customerrors

import "errors"

var ErrEmailAlreadyExists = errors.New("email already exist")
var ErrRecordNotFound = errors.New("record not found")
var ErrRecordAlreadyExist = errors.New("record already exist")
var ErrEmailInvalid = errors.New("invalid email or password")
var ErrPasswordInvalid = errors.New("invalid email or password")
var ErrNotVerified = errors.New("account not verified")

package model

import "errors"

var (
	ErrIDNotSpecified        = errors.New("user id not specified")
	ErrFirstNameNotSpecified = errors.New("firstname not specified")
	ErrLastNameNotSpecified  = errors.New("lastname not specified")
	ErrEmailNotSpecified     = errors.New("email not specified")
	ErrPasswordNotSpecified  = errors.New("password not specified")

	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

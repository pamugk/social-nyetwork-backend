package service

import (
	"errors"
)

var (
	NotFoundError         = errors.New("Not found")
	WrongCredentialsError = errors.New("Wrong credentials")
)

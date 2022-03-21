package user

import "errors"

var (
	ErrLoginExists = errors.New("login already exists")
	ErrAuth        = errors.New("invalid login or password")
)

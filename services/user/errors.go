package user

import "errors"

var (
	LoginExistsErr = errors.New("Login allready exists")
	AuthErr        = errors.New("Invalid login or password")
)

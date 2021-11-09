package errortype

import "errors"

var ErrResourceNotFound = errors.New("resource not found")
var ErrInvalidInput = errors.New("invalid input")
var ErrUnauthorized = errors.New("unauthorized credentials. wrong username or password")
var ErrUserNotFound = errors.New("user not found")

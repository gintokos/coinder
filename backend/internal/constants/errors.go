package constants

import "errors"

var (
	ErrServer = errors.New("server error")
	ErrParser = errors.New("parser error")
	ErrNotFound = errors.New("not found")
	ErrDatabase = errors.New("database error")
)

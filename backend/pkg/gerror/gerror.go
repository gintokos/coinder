package gerror

import (
	"errors"
	"fmt"
)

const (
	ErrCodeNotFound   = 404
	ErrCodeBadRequest = 400
	ErrCodeInternal   = 500
)

type GError struct {
	original      error
	tohighmessage error
	code          int
}

func FullError(err error) string {
	gerr, ok := err.(*GError)
	if !ok {
		return "error is not gerror"
	}
	orig := fmt.Sprintf("Original error: %s", gerr.original)
	thigh := fmt.Sprintf("To high level error: %s", gerr.tohighmessage)
	return fmt.Sprintf("%s ; %s", orig, thigh)
}

func (e *GError) Error() string {
	return e.tohighmessage.Error()
}

func (e *GError) Unwrap() error {
	return e.original
}

func New(original error, tohighmessage error, code int) *GError {
	if tohighmessage == nil {
		tohighmessage = errors.New("unknown error")
	}
	return &GError{
		original:      original,
		tohighmessage: tohighmessage,
		code:          code,
	}
}

func NewNotFound(original error, message error) *GError {
	return New(original, message, ErrCodeNotFound)
}

func NewBadRequest(original error, message error) *GError {
	return New(original, message, ErrCodeBadRequest)
}

func NewInternal(original error, message error) *GError {
	return New(original, message, ErrCodeInternal)
}

func Code(err error) int {
	gerr, ok := err.(*GError)
	if !ok {
		return -1
	}
	return gerr.code
}

func IsNotFound(err error) bool {
	gerr, ok := err.(*GError)
	if !ok {
		return false
	}
	return gerr.code == ErrCodeNotFound
}

func IsInvalidInput(err error) bool {
	gerr, ok := err.(*GError)
	if !ok {
		return false
	}
	return gerr.code == ErrCodeBadRequest
}

func IsInternal(err error) bool {
	gerr, ok := err.(*GError)
	if !ok {
		return false
	}
	return gerr.code == ErrCodeInternal
}

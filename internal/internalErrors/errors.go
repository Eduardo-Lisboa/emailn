package internalerrors

import "errors"

var (
	ErrInternal error = errors.New("Internal sever error")
)

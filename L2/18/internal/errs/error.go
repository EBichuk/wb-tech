package errs

import (
	"errors"
)

var (
	ErrNotFoundUser    = errors.New("not found user")
	ErrNotFoundEvent   = errors.New("not found event")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrInvalidDate     = errors.New("invalid date format")
	ErrInvalidPeriod   = errors.New("start time later than end time")
)

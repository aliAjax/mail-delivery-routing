package store

import "errors"

var ErrConflict = errors.New("conflict")
var ErrUnavailable = errors.New("store unavailable")

func IsRetryable(err error) bool { return errors.Is(err, ErrUnavailable) }

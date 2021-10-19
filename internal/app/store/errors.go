package store

import "errors"

var ErrRecordNotFound = errors.New("record not found")

var ErrConflict = errors.New("conflict")

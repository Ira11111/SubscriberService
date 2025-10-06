package repository

import "errors"

var (
	ErrFailedConnect = errors.New("Failed to connect database")
	ErrFailedClose   = errors.New("Failed to close database")
)

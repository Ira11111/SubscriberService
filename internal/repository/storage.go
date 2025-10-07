package repository

import "errors"

var (
	ErrFailedConnect = errors.New("Failed to connect database")
	ErrDataNotFoud   = errors.New("Data not found")
	ErrFailedSave    = errors.New("Failed to save data")
	ErrFailedGet     = errors.New("Failed to get data")
	ErrFailedDelete  = errors.New("Failed to delete data")
	ErrUpdateFailed  = errors.New("Failed to update data")
	ErrFailedScan    = errors.New("Failed to scan data")
)

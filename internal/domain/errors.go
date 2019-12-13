package domain

import "errors"

var (
	// ErrCouldNotSaveEntry represents a state where the given url/hash could not be saved
	ErrCouldNotSaveEntry = errors.New("could not save")
)

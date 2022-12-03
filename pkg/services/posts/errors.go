package posts

import (
	"errors"
)

var (
	// ErrPostNotFound ...
	ErrPostNotFound = errors.New("requested post could not be found")

	// ErrPostQuery ...
	ErrPostQuery = errors.New("requested posts could not be retrieved base on the given criteria")

	// ErrPostCreate ...
	ErrPostCreate = errors.New("post could not be created")

	// ErrPostUpdate ...
	ErrPostUpdate = errors.New("post could not be updated")
)

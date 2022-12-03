package users

import (
	"errors"
)

var (
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("requested user could not be found")

	// ErrUserQuery ...
	ErrUserQuery = errors.New("requested users could not be retrieved base on the given criteria")

	// ErrUserCreate ...
	ErrUserCreate = errors.New("user could not be created")

	// ErrUserUpdate ...
	ErrUserUpdate = errors.New("user could not be updated")

	// ErrUserDelete ...
	ErrUserDelete = errors.New("user could not be deleted")
)

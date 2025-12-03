package projectdomain

import "errors"

var (
	ErrProjectNotFound      error = errors.New("project not found")
	ErrProjectAlreadyExists error = errors.New("")
)

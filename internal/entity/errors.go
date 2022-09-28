package entity

import "errors"

var (
	ErrUserNotFound        = errors.New("user doesn't exists")
	ErrRepoNotFound        = errors.New("repo not found")
	ErrBookNotFound        = errors.New("book not found")
	ErrImageNotFound       = errors.New("image not found")
	ErrEmptyFile           = errors.New("file is empty")
	ErrUserAlreadyExists   = errors.New("user with such email already exists")
	ErrFileAlreadyExists   = errors.New("file with this name already exists")
	ErrUnknownCallbackType = errors.New("unknown callback type")
)

package models

import "errors"

var ErrNoRecord = errors.New("no record")
var ErrAlreadyExists = errors.New("already exists")
var ErrNoReference = errors.New("no reference")
var ErrInvalidAPI = errors.New("invalid api name")

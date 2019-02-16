package model

import "errors"

var (
	ERROR_USER_NOTEXISTS = errors.New("user not exist")
	ERROR_USER_EXISTS    = errors.New("user exist")
	ERROR_USER_PWD       = errors.New("pwd wrong")
)

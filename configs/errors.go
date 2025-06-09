package configs

import (
	"errors"
)

var (
	ErrNoConfigPath = errors.New("CONFIG_PATH not set")
	ErrIsNotExist = errors.New("config file is not exist")
	ErrReadFail = errors.New("cannot read the config file")
)

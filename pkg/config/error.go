package config

import "errors"

var (
	errInvalidDomain   = errors.New("Config error: inavalid domain URL")
	errMissingDBConfig = errors.New("Config error: missing database configuration")
)

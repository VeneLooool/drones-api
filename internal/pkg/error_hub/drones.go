package error_hub

import "errors"

var (
	ErrDroneNotFound = errors.New("drone not found")
)

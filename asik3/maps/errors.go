package maps

import "errors"

var (
	ErrUpstream        = errors.New("upstream error")
	ErrNegativeDistance = errors.New("negative distance from upstream")
)
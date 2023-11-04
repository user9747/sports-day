package cache

import (
	"errors"
	"time"

	redis "github.com/go-redis/redis/v8"
)

const suffix = "-lendingapp"

var (
	Nil                   = redis.Nil
	ErrorEmptyKey         = errors.New("redis: key cannot be blank string")
	ErrorEmptyPattern     = errors.New("redis: pattern cannot be blank string")
	ErrorUnsupportedValue = errors.New("redis: unsupported value passed")
)

const (
	contextTimeout = 20 * time.Second
)

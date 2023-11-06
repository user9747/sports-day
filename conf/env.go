package conf

import (
	"os"
)

var MODE = os.Getenv("MODE")

const (
	MODE_LOCAL_DOCKER = "local-docker"
)

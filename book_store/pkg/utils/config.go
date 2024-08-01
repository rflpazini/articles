package utils

import "os"

const (
	DEV  = "dev"
	STAG = "staging"
	PROD = "production"
)

func GetEnv() string {
	env := os.Getenv("ENV")
	if env == PROD {
		return PROD
	} else if env == STAG {
		return STAG
	} else {
		return DEV
	}
}

func GetHostName() string {
	name, _ := os.Hostname()
	if os.Getenv("STACK") != "" {
		name = os.Getenv("STACK")
	}

	return name
}

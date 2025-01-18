package env

import (
	"os"
	"strconv"
)

func GetInt(variable string, fallback int) int {
	result, err := strconv.Atoi(os.Getenv(variable))
	if err != nil {
		return fallback
	}
	return result
}

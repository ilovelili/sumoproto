package position

import (
	"os"
)

// Getwd get working directory
func Getwd() string {
	wd, _ := os.Getwd()
	return wd
}

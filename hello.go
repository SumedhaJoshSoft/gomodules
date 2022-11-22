package hello

import (
	"fmt"

	"github.com/fatih/color"
)

func Hello(name string) string {
	color.Cyan("Welcome %s", name)
	msg := fmt.Sprint("Hi, %s. Welcome!", name)
	return msg
}

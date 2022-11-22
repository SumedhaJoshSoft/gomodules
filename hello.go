package hello

import (
	"fmt"

	"github.com/fatih/color"
)

func Hello(name string) string {
	color.Cyan("Welcome %s", name)
	msg := fmt.Sprintf("Hi, %s. Welcome!", name)
	return msg
}

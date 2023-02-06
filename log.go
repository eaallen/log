package log

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func Warn(text ...string) {
	color.Yellow(strings.Join(text[0:], " "))
}

func Info(text ...string) {
	color.Cyan(strings.Join(text[0:], " "))
}

func Success(text ...string) {
	color.Green(strings.Join(text[0:], " "))
}

func Error(text ...string) {
	color.Red(strings.Join(text[0:], " "))
}

func Out(text ...string) {
	fmt.Println(strings.Join(text[0:], " "))
}

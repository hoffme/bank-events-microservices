package logger

import (
	"fmt"
	"log"
)

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorYellow = "\033[33m"
const colorPurple = "\033[35m"
const colorGray = "\033[37m"

var methodColor = map[string]string{
	"DEBUG ": colorGreen,
	"INFO ":  colorGray,
	"WARN ":  colorYellow,
	"ERROR":  colorRed,
	"FATAL":  colorPurple,
}

func send(method string, template string, extra ...any) {
	color := methodColor[method]
	log.Printf("%s%s %s%s", color, method, fmt.Sprintf(template, extra...), colorReset)
}

func Debug(template string, data ...any) {
	send("DEBUG", template, data...)
}

func Info(template string, data ...any) {
	send("INFO ", template, data...)
}

func Warn(template string, data ...any) {
	send("WARN ", template, data...)
}

func Error(template string, data ...any) {
	send("ERROR", template, data...)
}

func Fatal(template string, data ...any) {
	send("FATAL", template, data...)
}

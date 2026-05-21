package main

import "fmt"

type Formatter interface {
	Format(level, message string) string
}

type JSONFormatter struct{}

func (f JSONFormatter) Format(level, message string) string {
	text := fmt.Sprintf("{\"level\":\"%v\", \"message\":\"%v\"}", level, message)
	return text
}

type PlainTextFormatter struct{}

func (f PlainTextFormatter) Format(level, message string) string {
	text := fmt.Sprintf("[%v]: %v", level, message)
	return text

}

type Logger struct {
	formatter Formatter
}

type InfoLogger struct {
	Logger
}

func (l InfoLogger) Log(message string) {
	res := l.formatter.Format("INFO", message)
	fmt.Println(res)
}

func NewInfoLogger(f Formatter) *InfoLogger {
	return &InfoLogger{Logger{formatter: f}}
}

type ErrorLogger struct {
	Logger
}

func (l ErrorLogger) Log(message string) {
	res := l.formatter.Format("ERROR", message)
	fmt.Println(res)
}

func NewErrorLogger(f Formatter) *ErrorLogger {
	return &ErrorLogger{Logger{formatter: f}}
}
func main() {
	infoLogger := NewInfoLogger(JSONFormatter{})
	errLogger := NewErrorLogger(PlainTextFormatter{})

	infoLogger.Log("meow1")
	infoLogger.Log("meow2")

	errLogger.Log("meow3")
	errLogger.formatter = JSONFormatter{}
	errLogger.Log("meow4")
}

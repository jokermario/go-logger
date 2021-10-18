package src

import (
	"fmt"
	"time"
)

//go:generate stringer -type=LogAttribute
type LogAttribute int

const (
	Info LogAttribute = 1 << iota
	Warn
	Error
)

type Logger struct {
	timeFormat string
	debug      bool
}

func New(timeFormat string, debug bool) *Logger {
	return &Logger{
		timeFormat: timeFormat,
		debug:      debug,
	}
}

func (l *Logger) Debug(b bool) {
	l.debug = b
}

func (l *Logger) Log(level LogAttribute, statement string) {
	switch level {
	case Info, Warn:
		if l.debug {
			l.write(level, statement)
		}
	default:
		l.write(level, statement)
	}

	// fmt.Printf("%s %s\n", time.Now().Format(l.timeFormat), statement)
}

func (l *Logger) write(level LogAttribute, s string) {
	switch level {
	case Info:
		fmt.Printf("\033[34m[%s] %s %s\n", level, time.Now().Format(l.timeFormat), s)
	case Warn:
		fmt.Printf("\033[33m[%s] %s %s\n", level, time.Now().Format(l.timeFormat), s)
	case Error:
		fmt.Printf("\033[31m[%s] %s %s\n", level, time.Now().Format(l.timeFormat), s)
	default:
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		panic("invalid log level")
	}
}

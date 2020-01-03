package sudoku

import (
	"bytes"
	"log"
)

type Option func(l *Logger)

type Logger struct {
	debug bool
}

func NewLogger(options ...Option) *Logger {
	logger := &Logger{}
	for _, option := range options {
		option(logger)
	}
	return logger
}

func (l *Logger) Debug(s string, f ...interface{}) {
	if !l.debug {
		return
	}

	var b bytes.Buffer
	for i := 0; i < guessesDeep; i++ {
		b.WriteString("\t")
	}
	// fmt.Printf(b.String())
	// if !strings.HasSuffix(s, "\n") {
	// 	s = s + "\n"
	// }
	s = b.String() + s
	log.Printf(s, f...)
}

func OptionDebug(l *Logger) {
	l.debug = true
}

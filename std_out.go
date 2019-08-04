package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	NAME_LENGTH = 8
)

type StdOut struct {
	Logger *log.Logger
	flag   Severity
}

func (s *StdOut) Enable(prefix string, flags ...Severity) Log {
	flag := BuildFlag(flags...)
	return &StdOut{
		log.New(os.Stdout, fmt.Sprintf("[ %v ] ", padRight(prefix, " ", NAME_LENGTH)), log.LstdFlags|log.Lmicroseconds),
		flag,
	}
}

func (s *StdOut) Disable() Log {
	s.Logger = nil
	return s
}

func (s *StdOut) print(severity Severity, format string, params ...interface{}) {
	if s.Logger == nil {
		return
	}

	if s.flag&severity == 0 {
		return
	}
	severityText := "#"
	switch severity {
	case Warning:
		severityText = "W"
	case Error:
		severityText = "E"
	case Message:
		severityText = "M"
	case Debug:
		severityText = "D"
	}

	s.Logger.Printf(fmt.Sprintf("[%v] %v", severityText, format), params...)
}

func (s *StdOut) Debug(format string, params ...interface{}) {
	s.print(Debug, format, params...)
}

func (s *StdOut) Message(format string, params ...interface{}) {
	s.print(Message, format, params...)
}

func (s *StdOut) Warning(format string, params ...interface{}) {
	s.print(Warning, format, params...)
}

func (s *StdOut) Error(format string, params ...interface{}) {
	s.print(Error, format, params...)
}

func (s *StdOut) Panic(format string, err error, params map[string]interface{}) {
	str := fmt.Sprintf(format)
	if s.Logger != nil {
		s.Logger.Output(3, str)
		s.Logger.Printf("%v", err)
	}

	//panic(str)
}

func padRight(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

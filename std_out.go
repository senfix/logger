package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	NAME_LENGTH = 4
)

type StdOut struct {
	Logger *log.Logger
}

func (s *StdOut) Enable(prefix string) Log {
	s.Logger = log.New(os.Stdout, fmt.Sprintf("[ %v ] ", padRight(prefix, " ", NAME_LENGTH)), log.LstdFlags|log.Lmicroseconds)
	return s
}

func (s *StdOut) Disable() Log {
	s.Logger = nil
	return s
}

func (s *StdOut) Println(format string) {
	if s.Logger == nil {
		return
	}
	s.Logger.Println(format)
}

func (s *StdOut) Printf(format string, params ...interface{}) {
	if s.Logger == nil {
		return
	}
	s.Logger.Printf(format, params...)
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

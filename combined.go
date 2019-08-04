package logger

type Combined struct {
	Loggers []Log
}

func NewDevel() Log {
	return &Combined{
		Loggers: []Log{
			new(StdOut),
		},
	}
}

func NewProduction() Log {
	return &Combined{
		Loggers: []Log{
			new(StdOut),
			//NewSentry("https://837739f08aab4c21a9fb4402a36c6e0e:4e8f6af5b07144f58d6bc145580a07bc@sentry.io/1207685"),
		},
	}
}

func (s *Combined) Enable(prefix string, flags ...Severity) Log {
	n := &Combined{make([]Log, len(s.Loggers))}
	for i, l := range s.Loggers {
		n.Loggers[i] = l.Enable(prefix, flags...)
	}
	return n
}

func (s *Combined) Disable() Log {
	for _, l := range s.Loggers {
		l.Disable()
	}
	return s
}

func (s *Combined) Debug(format string, params ...interface{}) {
	for _, l := range s.Loggers {
		l.Debug(format, params...)
	}
}

func (s *Combined) Message(format string, params ...interface{}) {
	for _, l := range s.Loggers {
		l.Message(format, params...)
	}
}

func (s *Combined) Warning(format string, params ...interface{}) {
	for _, l := range s.Loggers {
		l.Message(format, params...)
	}
}

func (s *Combined) Error(format string, params ...interface{}) {
	for _, l := range s.Loggers {
		l.Error(format, params...)
	}
}

func (s *Combined) Panic(format string, err error, params map[string]interface{}) {
	for _, l := range s.Loggers {
		l.Panic(format, err, params)
	}

}

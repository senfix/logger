package logger

type combined struct {
	Loggers []Log
}

func NewDevel() Log {
	return &combined{
		Loggers: []Log{
			new(StdOut),
		},
	}
}

func NewProduction() Log {
	return &combined{
		Loggers: []Log{
			new(StdOut),
			//NewSentry("https://837739f08aab4c21a9fb4402a36c6e0e:4e8f6af5b07144f58d6bc145580a07bc@sentry.io/1207685"),
		},
	}
}

func (s *combined) Enable(prefix string) Log {
	for _, l := range s.Loggers {
		l.Enable(prefix)
	}
	return s
}

func (s *combined) Disable() Log {
	for _, l := range s.Loggers {
		l.Disable()
	}
	return s
}

func (s *combined) Println(format string) {
	for _, l := range s.Loggers {
		l.Println(format)
	}
}

func (s *combined) Printf(format string, params ...interface{}) {
	for _, l := range s.Loggers {
		l.Printf(format, params...)
	}
}

func (s *combined) Panic(format string, err error, params map[string]interface{}) {
	for _, l := range s.Loggers {
		l.Panic(format, err, params)
	}
}

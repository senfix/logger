package logger

type Severity int

const (
	Error Severity = 1 << iota
	Warning
	Message
	Debug
)

const (
	Low      = Error | Warning
	Default  = Low | Message
	Extended = Default | Debug
)

type Log interface {
	Enable(prefix string, flags ...Severity) Log
	Disable() Log
	Debug(format string, params ...interface{})
	Message(format string, params ...interface{})
	Warning(format string, params ...interface{})
	Error(format string, params ...interface{})
	Panic(message string, err error, params map[string]interface{})
}

func BuildFlag(flags ...Severity) (flag Severity) {
	if len(flags) <= 0 {
		flag = Default
	} else {
		for _, s := range flags {
			flag = flag | s
		}
	}
	return
}

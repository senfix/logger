package logger

type Log interface {
	Enable(prefix string) Log
	Disable() Log
	Println(format string)
	Printf(format string, params ...interface{})
	Panic(message string, err error, params map[string]interface{})
}

package hero

import "log"

// Logger is an interface for logging
type Logger interface {
	Print(v ...interface{})
	PrintLn(v ...interface{})
	Printf(fmt string, v ...interface{})
}

// DefaultLogger implements a Logger interface that writes logs to stdout
type DefaultLogger struct {
	log *log.Logger
}

//Print prints v into stdout
func (l *DefaultLogger) Print(v ...interface{}) {
	l.log.Print(v...)
}

//Println prints v in a new line
func (l *DefaultLogger) Println(v ...interface{}) {
	l.log.Println(v...)
}

//Printf print formatsp
func (l *DefaultLogger) Printf(fmt string, v ...interface{}) {
	l.log.Printf(fmt, v...)
}

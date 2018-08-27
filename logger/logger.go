package logger

import (
	"log"
	"os"
	"runtime"
)

//--------------------
// LOG LEVEL
//--------------------

// Log levels to control the logging output.
const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

// logLevel controls the global log level used by the logger.
var level = LevelTrace

// LogLevel returns the global log level and can be used in
// own implementations of the logger interface.
func Level() int {
	return level
}

// SetLogLevel sets the global log level used by the simple
// logger.
func SetLevel(l int) {
	level = l
}

// logger references the used application logger.
var BeeLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

// SetLogger sets a new logger.
func SetLogger(l *log.Logger) {
	BeeLogger = l
}

// Trace logs a message at trace level.
func Trace(v ...interface{}) {
	if level <= LevelTrace {
		BeeLogger.Printf("[T] %v\n", v)
	}
}

// Debug logs a message at debug level.
func Debug(v ...interface{}) {
	if level <= LevelDebug {
		BeeLogger.Printf("[D] %v\n", v)
	}
}

// Info logs a message at info level.
func Info(v ...interface{}) {
	if level <= LevelInfo {
		BeeLogger.Printf("[I] %v\n", v)
	}
}

// Warning logs a message at warning level.
func Warn(v ...interface{}) {
	if level <= LevelWarning {
		BeeLogger.Printf("[W] %v\n", v)
	}
}

// Error logs a message at error level.
func Error(v ...interface{}) {
	if level <= LevelError {
		pc, file, line, _ := runtime.Caller(1)
		//		log.Println(pc)
		//		log.Println(file)
		//		log.Println(line)
		//		log.Println(ok)
		f := runtime.FuncForPC(pc)
		//log.Println(f.Name())

		BeeLogger.Printf("[E] %v %v %v %v\n", v, file, line, f.Name())
	}
}

// Critical logs a message at critical level.
func Critical(v ...interface{}) {
	if level <= LevelCritical {
		BeeLogger.Printf("[C] %v\n", v)
	}
}

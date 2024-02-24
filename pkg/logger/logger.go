package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

var (
	logger *Logger
)

func GetLogger(prefix string) *Logger {
	logger = newLogger(prefix)
	return logger
}

func newLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

func (logger *Logger) Debug(v ...interface{}) {
	logger.debug.Print(v...)
}
func (logger *Logger) DebugWithValues(format string, v ...interface{}) {
	logger.debug.Printf(format, v...)
}

func (logger *Logger) Info(v ...interface{}) {
	logger.info.Print(v...)
}
func (logger *Logger) InfoWithValues(format string, v ...interface{}) {
	logger.info.Printf(format, v...)
}

func (logger *Logger) Warn(v ...interface{}) {
	logger.warning.Print(v...)
}
func (logger *Logger) WarnWithValues(format string, v ...interface{}) {
	logger.warning.Printf(format, v...)
}

func (logger *Logger) Error(v ...interface{}) {
	logger.err.Print(v...)
}
func (logger *Logger) ErrorWithValues(format string, v ...interface{}) {
	logger.err.Printf(format, v...)
}

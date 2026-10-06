package main

import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelInfo, LogLevelWarning, LogLevelError:
		return true
	default:
		return false
	}
}

type LogExtended struct {
	*log.Logger
	logLevel LogLevel // LogLevel это enum
}

func NewLogExtended() *LogExtended {
	return &LogExtended{
		Logger:   log.New(os.Stdout, "", log.LstdFlags),
		logLevel: LogLevelError,
	}
}

func (logger *LogExtended) SetLogLevel(l LogLevel) {
	if !l.IsValid() {
		return
	}

	logger.logLevel = l

}

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.logLevel < srcLogLvl {
		return
	}

	l.Logger.Println(prefix + msg)
}

func (l *LogExtended) Infoln(message string) {
	l.println(LogLevelInfo, "[INFO] ", message)
}

func (l *LogExtended) Errorln(message string) {
	l.println(LogLevelError, "[ERROR] ", message)
}

func (l *LogExtended) Warnln(message string) {
	l.println(LogLevelWarning, "[WARN] ", message)
}

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}

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

type LogExtended struct {
	*log.Logger
	logLevel LogLevel // LogLevel это enum
}

func NewLogExtended() LogExtended {
	lp := log.New(os.Stdout, "MY_EXT_LOGGER ", 3)
	return LogExtended{
		lp,
		LogLevelError,
	}
}

func (logger *LogExtended) SetLogLevel(l LogLevel) {
	logger.logLevel = l
}

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	// игнорируем сообщения, если уровень логгера меньше scrLogLvl
	// ...
	if l.logLevel >= srcLogLvl {
		l.Logger.Println(prefix + msg)
	}
}

func (l *LogExtended) Infoln(i string) {
	l.println(LogLevelInfo, "[INFO] ", i)
}

func (l *LogExtended) Errorln(i string) {
	l.println(LogLevelError, "[ERROR] ", i)
}

func (l *LogExtended) Warnln(i string) {
	l.println(LogLevelWarning, "[WARN] ", i)
}

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}

package main

import "log"

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
	return LogExtended{}
}

func (logger *LogExtended) SetLogLevel(l LogLevel) {
	logger.logLevel = l
}

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	// игнорируем сообщения, если уровень логгера меньше scrLogLvl
	// ...

	l.Logger.Println(prefix + msg)
}

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}

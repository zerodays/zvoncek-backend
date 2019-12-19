package logger

import (
	"fmt"
	"log"
)

type Level int

const (
	LevelDebug   Level = 0
	LevelInfo    Level = 1
	LevelWarning Level = 2
	LevelFatal   Level = 3
)

var MinimumLevel = LevelInfo

func levelToString(level Level) string {
	switch level {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarning:
		return "WARNING"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func Log(msg, name string, level Level) {
	if level < MinimumLevel {
		return
	}

	msg = fmt.Sprintf("[%s :: %s] %s", levelToString(level), name, msg)

	if level == LevelFatal {
		log.Fatal(msg)
	} else {
		log.Println(msg)
	}
}

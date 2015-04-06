package iDlogger

import (
	"time"
)

type Event struct {
	Logger  *Logger
	Data    map[string]interface{}
	Time    time.Time
	Level   Level
	Message string
}

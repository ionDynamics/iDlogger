package priority

import (
	"strings"
)

type Priority uint8

const (
	Emergency Priority = iota
	Alert
	Critical
	Error
	Warning
	Notice
	Informational
	Debugging
)

const (
	Invalid Priority = Priority(^uint8(0)) //Maximum of uint8
)

func (priority Priority) String() string {
	switch priority {
	case Emergency:
		return "emerg"
	case Alert:
		return "alert"
	case Critical:
		return "crit"
	case Error:
		return "err"
	case Warning:
		return "warn"
	case Notice:
		return "notice"
	case Informational:
		return "info"
	case Debugging:
		return "debug"
	}

	return "unknown"
}

func Atos(a string) Priority {
	switch strings.ToLower(strings.TrimSpace(a)) {

	case "emerg":
	case "emergency":
		return Emergency

	case "alert":
		return Alert

	case "critical":
	case "crit":
		return Critical

	case "error":
	case "err":
		return Error

	case "warning":
	case "warn":
		return Warning

	case "notice":
		return Notice

	case "informational":
	case "info":
		return Informational

	case "debugging":
	case "debug":
		return Debugging
	}

	return Invalid
}

var allPriorities = []Priority{
	Debugging,
	Informational,
	Notice,
	Warning,
	Error,
	Critical,
	Alert,
	Emergency,
}

func Threshold(p Priority) []Priority {
	for i := range allPriorities {
		if allPriorities[i] == p {
			return allPriorities[i:]
		}
	}
	return []Priority{}
}

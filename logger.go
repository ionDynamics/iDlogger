package iDlogger

import (
	"os"
	"sync"
	"time"
)

type Logger struct {
	Async       bool
	mu          sync.RWMutex
	prefix      string
	flag        uint8
	errCallback func(error)
	levelHooks  map[Level][]Hook
}

func New() *Logger {
	var sf *StdFormatter
	stdOut := new(StdHook)
	stdErr := new(StdHook)

	stdOut.SetFormatter(sf)
	stdErr.SetFormatter(sf)

	stdOut.SetWriter(os.Stdout)
	stdErr.SetWriter(os.Stderr)

	stdOut.SetLevels([]Level{
		DebugLevel,
		InfoLevel,
		WarnLevel,
	})

	stdErr.SetLevels([]Level{
		ErrorLevel,
		FatalLevel,
		PanicLevel,
	})

	log := &Logger{
		false,
		sync.RWMutex{},
		"",
		0,
		nil,
		map[Level][]Hook{
			PanicLevel: []Hook{},
			FatalLevel: []Hook{},
			ErrorLevel: []Hook{},
			WarnLevel:  []Hook{},
			InfoLevel:  []Hook{},
			DebugLevel: []Hook{},
		},
	}

	log.AddHook(stdOut)
	log.AddHook(stdErr)

	return log
}

func (l *Logger) dispatch(e *Event) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for _, h := range l.levelHooks[e.Level] {
		err := h.Fire(e)
		if err != nil && l.errCallback != nil {
			l.errCallback(err)
		}
	}
}

func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l *Logger) SetErrCallback(errCallback func(error)) {
	l.errCallback = errCallback
}

func (l *Logger) AddHook(h Hook) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, lvl := range h.Levels() {
		l.levelHooks[lvl] = append(l.levelHooks[lvl], h)
	}
}

func (l *Logger) Log(e *Event) {
	if l.Async {
		go l.dispatch(e)
	} else {
		l.dispatch(e)
	}
}

func (l *Logger) Debug(entry string) {
	l.Log(&Event{l, map[string]interface{}{}, time.Now(), DebugLevel, entry})
}

func (l *Logger) Info(entry string) {
	l.Log(&Event{l, map[string]interface{}{}, time.Now(), InfoLevel, entry})
}

func (l *Logger) Warn(entry string) {
	l.Log(&Event{l, map[string]interface{}{}, time.Now(), WarnLevel, entry})
}

func (l *Logger) Error(entry string) {
	l.Log(&Event{l, map[string]interface{}{}, time.Now(), ErrorLevel, entry})
}

func (l *Logger) Fatal(entry string) {
	l.Log(&Event{l, map[string]interface{}{}, time.Now(), FatalLevel, entry})
}

func (l *Logger) Panic(entry string) {
	l.Log(&Event{l, map[string]interface{}{}, time.Now(), PanicLevel, entry})
}

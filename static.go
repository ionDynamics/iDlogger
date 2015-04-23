package iDlogger

import (
	"fmt"
	"os"
	"time"
)

var (
	std = New()
)

func StandardLogger() *Logger {
	return std
}

func Wait() {
	std.wg.Wait()
}

func SetPrefix(prefix string) {
	std.prefix = prefix
}

func SetErrCallback(errCallback func(error)) {
	std.errCallback = errCallback
}

func AddHook(h Hook) {
	std.mu.Lock()
	defer std.mu.Unlock()

	for _, lvl := range h.Levels() {
		std.levelHooks[lvl] = append(std.levelHooks[lvl], h)
	}
}

func Log(e *Event) {
	std.wg.Add(1)
	if std.Async {
		go std.dispatch(e)
	} else {
		std.dispatch(e)
		std.wg.Wait()
	}
}

func Debug(entry ...interface{}) {
	std.Log(&Event{std, map[string]interface{}{}, time.Now(), DebugLevel, fmt.Sprint(entry...)})
}

func Info(entry ...interface{}) {
	std.Log(&Event{std, map[string]interface{}{}, time.Now(), InfoLevel, fmt.Sprint(entry...)})
}

func Warn(entry ...interface{}) {
	std.Log(&Event{std, map[string]interface{}{}, time.Now(), WarnLevel, fmt.Sprint(entry...)})
}

func Error(entry ...interface{}) {
	std.Log(&Event{std, map[string]interface{}{}, time.Now(), ErrorLevel, fmt.Sprint(entry...)})
}

func Fatal(entry ...interface{}) {
	std.Log(&Event{std, map[string]interface{}{}, time.Now(), FatalLevel, fmt.Sprint(entry...)})
	std.Wait()
	os.Exit(1)
}

func Panic(entry ...interface{}) {
	std.Log(&Event{std, map[string]interface{}{}, time.Now(), PanicLevel, fmt.Sprint(entry...)})
	std.Wait()
	panic(fmt.Sprint(entry...))
}

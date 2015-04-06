package iDlogger

import (
	"io"
)

type Hook interface {
	Levels() []Level
	Fire(*Event) error
}

type StdHook struct {
	l []Level
	w io.Writer
	f Formatter
}

func (sh *StdHook) Fire(e *Event) error {
	byt, err := sh.f.Format(e)
	if err == nil {
		_, err = sh.w.Write(*byt)
	}
	return err
}

func (sh *StdHook) Levels() []Level {
	return sh.l
}

func (sh *StdHook) SetLevels(l []Level) {
	sh.l = l
}

func (sh *StdHook) SetWriter(w io.Writer) {
	sh.w = w
}

func (sh *StdHook) SetFormatter(f Formatter) {
	sh.f = f
}

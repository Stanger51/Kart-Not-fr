package context

import (
	"bytes"
	"fmt"
	"io"

	"github.com/kartFr/Asset-Reuploader/internal/color"
)

type logger struct {
	History *bytes.Buffer
	writer  io.Writer
}

func newLogger() *logger {
	var b bytes.Buffer
	w := io.MultiWriter(
		&b,
		color.Output,
	)

	return &logger{
		History: &b,
		writer:  w,
	}
}

func (l *logger) Error(a ...any) {
	color.Error.Fprintln(l.writer, a...)
}

func (l *logger) Info(a ...any) {
	color.Info.Fprintln(l.writer, a...)
}

func (l *logger) Println(a ...any) {
	fmt.Fprintln(l.writer, a...)
}

func (l *logger) Success(a ...any) {
	color.Success.Fprintln(l.writer, a...)
}

func (l *logger) Warn(a ...any) {
	color.Warn.Fprintln(l.writer, a...)
}

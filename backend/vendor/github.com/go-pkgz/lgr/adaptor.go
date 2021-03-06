package lgr

import (
	"log"
	"strings"
)

// Writer holds lgr.L and wraps with io.Writer interface
type Writer struct {
	L
	level string // if defined added to each message
}

// Write to lgr.L
func (w *Writer) Write(p []byte) (n int, err error) {
	w.Logf(w.level + string(p))
	return len(p), nil
}

// ToWriter makes io.Writer for given lgr.L with optional level
func ToWriter(l L, level string) *Writer {
	if level != "" && !strings.HasSuffix(level, " ") {
		level += " "
	}
	return &Writer{l, level}
}

func ToStdLogger(l L, level string) *log.Logger {
	return log.New(ToWriter(l, level), "", 0)
}

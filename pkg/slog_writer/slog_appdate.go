package slog_writer

import "log/slog"

type SlogWriter struct {
	Logger *slog.Logger
}

func (w *SlogWriter) Write(p []byte) (n int, err error) {
	w.Logger.Info(string(p))
	return len(p), nil
}

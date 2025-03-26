package slogwriter

import "log/slog"

type SlogWriter struct {
	Logger *slog.Logger
}

func (w *SlogWriter) Write(p []byte) (int, error) {
	w.Logger.Info(string(p))
	return len(p), nil
}

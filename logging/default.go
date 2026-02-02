package logging

import "log/slog"

// TODO: Set default logging level based on environment
var DefaultLevel = slog.LevelDebug

func init() {
	slog.SetDefault(NewConsoleLogger(DefaultLevel))
}

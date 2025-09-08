package logging

import (
	"context"
	"log/slog"
	"os"
)

// 定义两个更高优先级的日志级别
const (
	LevelPanic = slog.Level(12)
	LevelFatal = slog.Level(16)
)

// Config 定义了日志系统的配置
type Config struct {
	Level     string `yaml:"level"`
	Format    string `yaml:"format"`
	AddSource bool   `yaml:"add_source"`
}

var globalLogger *slog.Logger

// Init 根据配置初始化全局 logger
func Init(cfg Config) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		AddSource: cfg.AddSource,
		Level:     level,
	}

	var handler slog.Handler
	switch cfg.Format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	case "console":
		fallthrough
	default:
		handler = NewConsoleHandler(os.Stdout, opts)
	}

	globalLogger = slog.New(handler)
	slog.SetDefault(globalLogger)
}

// Get 返回全局 logger 实例
func Get() *slog.Logger {
	if globalLogger == nil {
		Init(Config{Level: "info", Format: "console", AddSource: false})
	}
	return globalLogger
}

// Panic 记录一条 PANIC 级别的日志，然后调用 panic()。
func Panic(msg string, args ...any) {
	Get().Log(context.Background(), LevelPanic, msg, args...)
	panic(buildPanicMessage(msg))
}

// Fatal 记录一条 FATAL 级别的日志，然后调用 os.Exit(1)。
func Fatal(msg string, args ...any) {
	Get().Log(context.Background(), LevelFatal, msg, args...)
	os.Exit(1)
}

// 辅助函数，用于从 slog 参数构建一个更丰富的 panic 消息
func buildPanicMessage(msg string) string {
	panicMsg := msg
	return panicMsg
}

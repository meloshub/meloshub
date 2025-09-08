package logging

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"unicode"
)

// ANSI 颜色代码
const (
	colorPurple = "\033[35m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorGray   = "\033[90m"
	colorReset  = "\033[0m"
)

// ConsoleHandler 实现了 slog.Handler 接口，提供彩色、人类可读的日志输出。
type ConsoleHandler struct {
	opts   slog.HandlerOptions
	writer io.Writer
	mu     *sync.Mutex
	attrs  []slog.Attr
	groups []string
}

// NewConsoleHandler 创建一个新的 ConsoleHandler
func NewConsoleHandler(w io.Writer, opts *slog.HandlerOptions) *ConsoleHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	return &ConsoleHandler{
		opts:   *opts,
		writer: w,
		mu:     &sync.Mutex{},
	}
}

func (h *ConsoleHandler) Enabled(_ context.Context, level slog.Level) bool {
	minLevel := slog.LevelInfo
	if h.opts.Level != nil {
		minLevel = h.opts.Level.Level()
	}
	return level >= minLevel
}

func (h *ConsoleHandler) Handle(_ context.Context, r slog.Record) error {
	var buf []byte

	buf = r.Time.AppendFormat(buf, "[2006-01-02 15:04:05]")

	// 级别
	levelColor := ""
	levelText := r.Level.String()
	switch r.Level {
	case slog.LevelDebug:
		levelColor, levelText = colorGray, "DEBG"
	case slog.LevelInfo:
		levelColor, levelText = colorGreen, "INFO"
	case slog.LevelWarn:
		levelColor, levelText = colorYellow, "WARN"
	case slog.LevelError:
		levelColor, levelText = colorRed, "ERRO"
	case LevelPanic:
		levelColor, levelText = colorRed, "PANC"
	case LevelFatal:
		levelColor, levelText = colorPurple, "FATL"
	}
	buf = append(buf, ' ')
	buf = append(buf, levelColor...)
	buf = append(buf, levelText...)
	buf = append(buf, colorReset...)

	buf = append(buf, ' ')
	buf = append(buf, r.Message...)

	if h.opts.AddSource && r.Level >= slog.LevelWarn {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		frame, _ := fs.Next()
		buf = append(buf, ' ')
		buf = append(buf, colorGray...)
		buf = append(buf, '(')
		buf = append(buf, filepath.Base(frame.File)...)
		buf = append(buf, ':')
		buf = strconv.AppendInt(buf, int64(frame.Line), 10)
		buf = append(buf, ')')
		buf = append(buf, colorReset...)
	}

	attrs := h.collectAttrs(r)
	if len(attrs) > 0 {
		buf = append(buf, ' ')
	}
	for i, a := range attrs {
		h.appendAttr(&buf, a)
		if i < len(attrs)-1 {
			buf = append(buf, ' ')
		}
	}

	buf = append(buf, '\n')

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.writer.Write(buf)
	return err
}

func (h *ConsoleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandler := *h
	newHandler.attrs = append(newHandler.attrs, attrs...)
	return &newHandler
}

func (h *ConsoleHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	newHandler := *h
	newHandler.groups = append(newHandler.groups, name)
	return &newHandler
}

// collectAttrs 收集并组合 Handler 自身的属性和 Record 中的属性
func (h *ConsoleHandler) collectAttrs(r slog.Record) []slog.Attr {
	count := len(h.attrs)
	r.Attrs(func(slog.Attr) bool {
		count++
		return true
	})

	attrs := make([]slog.Attr, 0, count)
	attrs = append(attrs, h.attrs...)
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a)
		return true
	})

	// 将组应用于属性
	if len(h.groups) > 0 {
		prefix := ""
		for _, g := range h.groups {
			prefix += g + "."
		}
		for i := range attrs {
			attrs[i].Key = prefix + attrs[i].Key
		}
	}

	return attrs
}

// appendAttr 格式化单个属性
func (h *ConsoleHandler) appendAttr(buf *[]byte, a slog.Attr) {
	*buf = append(*buf, colorBlue...)
	*buf = append(*buf, a.Key...)
	*buf = append(*buf, colorReset...)
	*buf = append(*buf, '=')

	val := a.Value.String()
	if needsQuoting(val) {
		*buf = strconv.AppendQuote(*buf, val)
	} else {
		*buf = append(*buf, val...)
	}
}

// needsQuoting 检查字符串是否包含空格或非打印字符，以决定是否需要加引号
func needsQuoting(s string) bool {
	for _, r := range s {
		if unicode.IsSpace(r) || !unicode.IsPrint(r) {
			return true
		}
	}
	return false
}

package custom

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"strings"
)

type CustomHandler struct {
	slog.Handler
	w        io.Writer
	indent   int
	notifier Notifier
}

type Notifier interface {
	Notify(ctx context.Context, msg string) error
}

type CustomHandlerOptions struct {
	Indent   int
	Notifier Notifier
	SlogOpts slog.HandlerOptions
}

func NewCustomHandler(w io.Writer, opts CustomHandlerOptions) slog.Handler {
	return &CustomHandler{
		Handler:  slog.NewJSONHandler(w, &opts.SlogOpts),
		w:        w,
		indent:   opts.Indent,
		notifier: opts.Notifier,
	}
}

func (h *CustomHandler) Handle(ctx context.Context, rec slog.Record) error {
	requestId := getRequestID(ctx)
	rec.AddAttrs(slog.String("request_id", requestId))

	fields := make(map[string]any, rec.NumAttrs())
	fields[slog.TimeKey] = rec.Time
	fields[slog.LevelKey] = rec.Level
	fields[slog.MessageKey] = rec.Message

	rec.Attrs(func(a slog.Attr) bool {
		addFields(fields, a)
		return true
	})

	b, err := json.MarshalIndent(fields, "", strings.Repeat(" ", h.indent))
	if err != nil {
		return err
	}
	h.w.Write(b)

	if h.notifier != nil && rec.Level > slog.LevelWarn {
		if err := h.notifier.Notify(ctx, rec.Message); err != nil {
			return err
		}
	}

	return nil
}

func addFields(fields map[string]any, a slog.Attr) {
	value := a.Value.Any()
	if _, ok := value.([]slog.Attr); !ok {
		fields[a.Key] = value
		return
	}

	attrs := value.([]slog.Attr)
	// ネストしている場合、再起的にフィールドを探索する。
	innerFields := make(map[string]any, len(attrs))
	for _, attr := range attrs {
		addFields(innerFields, attr)
	}
	fields[a.Key] = innerFields
}

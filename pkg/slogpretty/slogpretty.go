package slogpretty

import (
	"fmt"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	stdLog "log"

	"github.com/fatih/color"
)

type PrettyHandlerOptions struct {
	SlogOpts *slog.HandlerOptions
}

type PrettyHandler struct {
	opts PrettyHandlerOptions
	slog.Handler
	l     *stdLog.Logger
	attrs []slog.Attr
}

func (opts PrettyHandlerOptions) NewPrettyHandler(
	out io.Writer,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       stdLog.New(out, "", 0),
	}

	return h
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := fmt.Sprintf("%5s:", r.Level.String())

	bgGreen := color.New(color.BgGreen, color.Bold)
	bgCyan := color.New(color.BgCyan, color.Bold)
	bgYellow := color.New(color.BgYellow, color.Bold)
	bgRed := color.New(color.BgRed, color.Bold)
	gray := color.RGB(127, 127, 127)

	switch r.Level {
	case slog.LevelDebug:
		level = bgGreen.Sprintf(level)
	case slog.LevelInfo:
		level = bgCyan.Sprintf(level)
	case slog.LevelWarn:
		level = bgYellow.Sprintf(level)
	case slog.LevelError:
		level = bgRed.Sprintf(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())

	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	var b []byte
	var err error

	if len(fields) > 0 {
		b, err = json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
	}

	timeStr := gray.Sprintf(r.Time.Format("[15:05:05.000]"))

	var msg string
	switch r.Level {
	case slog.LevelWarn:
		msg = color.YellowString(r.Message)
	case slog.LevelError:
		msg = color.RedString(r.Message)
	default:
		msg = r.Message
	}

	h.l.Println(
		timeStr,
		level,
		msg,
		color.WhiteString(string(b)),
	)

	return nil
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyHandler{
		Handler: h.Handler,
		l:       h.l,
		attrs:   attrs,
	}
}

func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	// TODO: implement
	return &PrettyHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}


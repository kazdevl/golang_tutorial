package main

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/kazdevl/golang_tutorial/logger/slog/custom"
)

var LevelNotice = slog.Level(2)

// TODO 独自のハンドラーの作成 & 独自のレベルの追加
func main() {
	var logger *slog.Logger
	logger = jsonLogger(false)
	logger.Debug("Debug message", slogAttr())
	logger.Info("Info message", slogAttr())
	logger.Warn("Warn message", slogAttr())
	logger.With(slog.String("logger", "json")).Warn("Warn message With `With` method", slogAttr())
	logger.WithGroup("jsonHandler").Warn("Warn message With `WithGroup` method", slogAttr())
	logger.Error("Error message", slogAttr())

	ctx := context.Background()
	logger.LogAttrs(ctx, slog.LevelInfo, "LogAttrs message", slogAttr())

	fmt.Println("********************")

	logger = textLogger(false)
	logger.Debug("Debug message", slogAttr())
	logger.Info("Info message", slogAttr())
	logger.Warn("Warn message", slogAttr())
	logger.With(slog.String("logger", "json")).Warn("Warn message With `With` method", slogAttr())
	logger.WithGroup("jsonHandler").Warn("Warn message With `WithGroup` method", slogAttr())
	logger.Error("Error message", slogAttr())

	fmt.Println("********************")
	logger = childLoggerPattern1(jsonLogger(false))
	logger.Info("Info Message from Child", slogAttr())

	fmt.Println("********************")
	logger = childLoggerPattern2(jsonLogger(false))
	logger.Info("Info Message from Child", slogAttr())

	fmt.Println("********************")
	logger = slog.New(slog.NewJSONHandler(os.Stdout, sanitizeHandlerOpt(false)))
	logger.Info("Info message",
		slog.String("password", "123456"),
		slog.String("username", "admin"),
		slog.Group("setting_content",
			slog.Bool("is_setting", true),
			slog.String("env", "local"),
			slog.Int("port", 8080),
		),
	)

	fmt.Println("********************")
	logger, close := fileoutputLogger("log.txt", slog.LevelWarn)
	defer close()

	logger.Info("Info message", slogAttr())
	logger.Warn("Warn message",
		slog.String("username", "admin"),
	)

	fmt.Println("********************")
	logger = textLogger(false)
	logger.Log(ctx, LevelNotice, "Notice message", slogAttr())
	logger = customLevelHandler()
	logger.Log(ctx, LevelNotice, "Notice message", slogAttr())

	fmt.Println("********************")
	notifier := &OSStdOutNotifier{
		b: new(bytes.Buffer),
	}
	handler := custom.NewCustomHandler(os.Stdout, custom.CustomHandlerOptions{
		Indent: 2,
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelInfo,
		},
		Notifier: notifier,
	})
	logger = slog.New(handler)

	ctx = custom.SetRequestID(ctx)
	logger.InfoContext(ctx, "Info message", slogAttr())
	logger.WarnContext(ctx, "Warn message", slogAttr())
	logger.ErrorContext(ctx, "Error message", slogAttr())

	fmt.Println("")
	fmt.Println("Buffer: ", notifier.String())
}

type OSStdOutNotifier struct {
	b *bytes.Buffer
}

func (n *OSStdOutNotifier) Notify(ctx context.Context, msg string) error {
	_, err := n.b.WriteString(msg)
	return err
}

func (n *OSStdOutNotifier) String() string {
	return n.b.String()
}

func jsonLogger(isEnableSource bool) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: isEnableSource,
		Level:     slog.LevelInfo,
	}))
	return logger
}

func textLogger(isEnableSource bool) *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: isEnableSource,
		Level:     slog.LevelInfo,
	}))
	return logger
}

func fileoutputLogger(fileName string, level slog.Level) (*slog.Logger, func()) {
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	opt := slog.NewJSONHandler(f, &slog.HandlerOptions{
		Level: &level,
	})

	logger := slog.New(opt)
	return logger, func() {
		f.Close()
	}
}

func customLevelHandler() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: LevelNotice,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key != slog.LevelKey {
				return a
			}
			if level := a.Value.Any().(slog.Level); level == LevelNotice {
				a.Value = slog.StringValue("NOTICE")
			}

			return a
		},
	}

	return slog.New(slog.NewTextHandler(os.Stdout, opts))
}

func sanitizeHandlerOpt(isEnableSource bool) *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: isEnableSource,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			fmt.Println("groups: ", groups)
			fmt.Println("attr: ", a)

			if a.Key == "password" {
				return slog.String("password", "*****")
			}
			return a
		},
	}
}

func childLoggerPattern1(logger *slog.Logger) *slog.Logger {
	return logger.With(slog.String("child", "child"))
}

func childLoggerPattern2(logger *slog.Logger) *slog.Logger {
	return logger.WithGroup("child")
}

func slogAttr() slog.Attr {
	v := slog.Group("setting_content", slog.Bool("is_setting", true), slog.String("env", "local"), slog.Int("port", 8080))
	// v := slog.String("default", "sample1")
	return v
}

package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// example1()
	// example2(os.Stdout)
	// f, err := os.Create("./logger/zap/log.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// example2(f)
	// example3()
	l, stopFn := exampleBuffer()
	defer stopFn()
	for i := 0; i < 1000; i++ {
		l.Info("Info message", zap.Int("count", i))
		time.Sleep(100 * time.Millisecond)
	}
}

func example1() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Info message",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func exampleBuffer() (*zap.Logger, func()) {
	ws := os.Stdout
	bws := &zapcore.BufferedWriteSyncer{WS: ws, FlushInterval: 5 * time.Second, Size: 10 * 1024}

	enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(enc, bws, zap.InfoLevel)

	logger := zap.New(core)
	return logger, func() {
		bws.Stop()
	}
}

// 参照: https://christina04.hatenablog.com/entry/golang-zap-tips
func example2(w io.Writer) {
	sink := zapcore.AddSync(w)
	lsink := zapcore.Lock(sink)

	enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(enc, lsink, zap.InfoLevel)

	logger := zap.New(core)
	logger.Info("Info message", zap.String("url", "http://example.com"), zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
}

func example3() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	now := time.Now()
	fmt.Println("Now:", now.Unix())

	date := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Date:", date.Unix())
	clock := constantClock(date)
	logger.WithOptions(zap.WithClock(clock)).Info("Info message with time")
}

type constantClock time.Time

func (c constantClock) Now() time.Time {
	return time.Time(c)
}

func (c constantClock) NewTicker(d time.Duration) *time.Ticker {
	return time.NewTicker(d)
}

var _ zapcore.Clock = constantClock{}

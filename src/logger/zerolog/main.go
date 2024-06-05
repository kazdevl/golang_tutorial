package main

import (
	"errors"
	"os"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	// example()
	// example2()
	// example_error()
	// example_logger()
	example_sampling()
}

func example() {
	zlog.Info().Str("key", "value").Msg("hello world")
}

func example2() {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	zlog.Info().Msg("hello world")
	zlog.Warn().Msg("hello world")
}

func example_error() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	err := outer()
	// NOTE: 機能していない
	zlog.Error().Stack().Err(err).Msg("outer error message")
}

func inner() error {
	return errors.New("inner error message")
}

func middle() error {
	return inner()
}

func outer() error {
	return middle()
}

func example_logger() {
	// これで出力先の調整をできはする
	f, err := os.Create("./logger/zerolog/log.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: f})
	zlog.Info().Msg("hello world in file")
}

func example_sampling() {
	sampled := zlog.Sample(&zerolog.BasicSampler{N: 5})
	for i := 0; i < 100; i++ {
		sampled.Info().Int("looCount", i).Msg("Hello world")
	}
}

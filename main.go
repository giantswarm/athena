package main

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/giantswarm/athena/cmd"
)

func main() {
	err := mainE(context.Background())
	if err != nil {
		panic(err)
	}
}

func mainE(ctx context.Context) error {
	var err error

	var log *zap.SugaredLogger
	{
		var logger *zap.Logger
		logger, err = zap.Config{
			Level:    zap.NewAtomicLevelAt(zap.DebugLevel),
			Encoding: "json",
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "time",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				FunctionKey:    "func",
				MessageKey:     "message",
				StacktraceKey:  zapcore.OmitKey,
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.RFC3339TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stdout"},
		}.Build()
		if err != nil {
			return microerror.Mask(err)
		}
		log = logger.Sugar().Named("f-logger")
	}

	var rootCommand *cobra.Command
	{
		c := cmd.Config{
			Log:   log,
			Viper: viper.New(),
		}

		rootCommand, err = cmd.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = rootCommand.ExecuteContext(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// MIT License
//
// Copyright (c) 2017 Lalamove.com
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
//
// Reference: https://lalamove.atlassian.net/wiki/spaces/TECH/pages/82149406/Kubernetes
// Lalamove kubernetes logging format
// {
// 		"message": "", // string describing what happened
// 		"src_file": "", // file path
// 		"src_line": "", // line number
// 		"fields": {}, // custom field here
// 		"level": "", // debug/info/warning/error/fatal
// 		"time": "", // ISO8601.nanoseconds+TZ (in node only support precision up to milliseconds)
// 		"backtrace": "" // err stack
// }
//
// Error: https://godoc.org/github.com/pkg/errors
package logs

import (
	"runtime"
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ISO8601 = "2006-01-02T15:04:05.000000000Z0700"

	Debug   = "debug"
	Info    = "info"
	Warning = "warning"
	Error   = "error"
	Fatal   = "fatal"
)

func NewLalamoveEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		CallerKey:      "src_file",
		MessageKey:     "message",
		StacktraceKey:  "backtrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    LalamoveLevelEncoder,
		EncodeTime:     LalamoveISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func LalamoveLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	if l.String() == zapcore.WarnLevel.String() {
		// Convert warn to warning
		enc.AppendString(Warning)
	} else {
		enc.AppendString(l.String())
	}
}

func LalamoveISO8601TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format(ISO8601))
}

func Logger() *zap.Logger {
	_, _, fl, _ := runtime.Caller(1)

	cfg := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    NewLalamoveEncoderConfig(),
		OutputPaths:      []string{"stdout", "/tmp/logs"},
		ErrorOutputPaths: []string{"stderr"},
	}

	showSourceLine := zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(
			c.With([]zapcore.Field{
				{
					Key:    "src_line",
					Type:   zapcore.StringType,
					String: strconv.Itoa(fl),
				},
			}),
		)
	})

	Logger, _ := cfg.Build()
	defer Logger.Sync()
	return Logger.WithOptions(showSourceLine)
}

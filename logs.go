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
	"encoding/json"
	"log"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	ISO8601 = "2006-01-02T15:04:05.000000000Z07:00"

	Debug   = "debug"
	Info    = "info"
	Warning = "warning"
	Error   = "error"
	Fatal   = "fatal"
)

var (
	Logger *zap.Logger

	configJSON = `{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase",
	    "timeKey": "time",
	    "stacktraceKey": "backtrace",
	    "callerKey": "src_file",
	    "lineEnding": "src_line"
	  }
	}`
)

func init() {
	var cfg zap.Config
	if err := json.Unmarshal([]byte(configJSON), &cfg); err != nil {
		log.Fatalln(errors.Wrap(err, "Fail to init zap logger."))
	}
	Logger, err := cfg.Build()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Invalid configJSON for the zap logger."))
	}
	Logger.Info("Logger started")
	defer Logger.Sync()
}

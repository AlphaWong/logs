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
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const (
	ISO8601 = "2006-01-02T15:04:05.000000000Z07:00"

	Debug   = "debug"
	Info    = "info"
	Warning = "warning"
	Error   = "error"
	Fatal   = "fatal"
)

type LalamoveLog struct {
	Message    string            `json:"message"`
	SourceFile string            `json:"src_file"`
	SourceLine string            `json:"src_line"`
	Fields     map[string]string `json:"fields"`
	Level      string            `json:"level"`
	Time       string            `json:"time"`
	Backtrace  string            `json:"backtrace"`
}

// String will convert LalamoveLog to json string
// return a string
func (llmlog *LalamoveLog) String() (s string) {
	var b bytes.Buffer
	json.NewEncoder(&b).Encode(llmlog)
	s = b.String()
	return
}

// ErrorLogging will log the error message
// err error is the error struct from the source file
// l string is the error level ( e.g debug, error, warning etc )
// cf map[string]string is the custom fields offer the extra information of the error
func FireLalamoveLog(err error, l string, cf map[string]string) {
	if err != nil {
		_, fn, ln, _ := runtime.Caller(1)
		logMsg := NewLalamoveLog(
			err,
			fn,
			strconv.Itoa(ln),
			cf,
			l,
		)
		log.Println(logMsg.String())
	}
}

// NewLalamoveLog will create a lalamove log struct based on the error, source file name, source file line number, custom fields and log level
// err error is the error struct from the source file
// fn string is the source file name
// ln string is the source file line number
// cf map[string]string is the custom fields offer the extra information of the error
// l string is the error level ( e.g debug, error, warning etc )
func NewLalamoveLog(err error, fn string, ln string, cf map[string]string, l string) *LalamoveLog {
	return &LalamoveLog{
		Message:    err.Error(),
		SourceFile: fn,
		SourceLine: ln,
		Fields:     cf,
		Level:      l,
		Time:       time.Now().UTC().Format(ISO8601),
		Backtrace:  GetErrorStackTrace(err),
	}
}

// GetErrorStackTrace will get the error stack trace then convert it to string
// ccause error is the original error
// return a string contains the error stack trace
func GetErrorStackTrace(cause error) (s string) {
	err := errors.WithStack(cause)
	return fmt.Sprintf("%+v", err)
}

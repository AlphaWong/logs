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

package logs

import (
	"encoding/json"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func generateTestingError() (err error, fn string, ln string, cf map[string]string, l string) {
	err = errors.New("I am a testing error")
	_, fn, lineNumber, _ := runtime.Caller(1)
	ln = strconv.Itoa(lineNumber)
	cf = make(map[string]string, 0)
	l = Debug
	return
}

func TestNewLalamoveLogPass(t *testing.T) {
	err, fn, lineNumber, cf, Debug := generateTestingError()

	actualLalamoveLog := NewLalamoveLog(err, fn, lineNumber, cf, Debug)

	assert.Equal(t, err.Error(), actualLalamoveLog.Message)
	assert.Equal(t, fn, actualLalamoveLog.SourceFile)
	assert.Equal(t, lineNumber, actualLalamoveLog.SourceLine)
	assert.Equal(t, cf, actualLalamoveLog.Fields)
	assert.Equal(t, Debug, actualLalamoveLog.Level)
}

func TestLalamoveLogJSONPass(t *testing.T) {
	err, fn, lineNumber, cf, Debug := generateTestingError()

	actualLalamoveLog := NewLalamoveLog(err, fn, lineNumber, cf, Debug)
	var expectedLalamoveLog LalamoveLog

	json.NewDecoder(strings.NewReader(actualLalamoveLog.String())).Decode(&expectedLalamoveLog)

	assert.Equal(t, expectedLalamoveLog.Message, actualLalamoveLog.Message)
	assert.Equal(t, expectedLalamoveLog.SourceFile, actualLalamoveLog.SourceFile)
	assert.Equal(t, expectedLalamoveLog.SourceLine, actualLalamoveLog.SourceLine)
	assert.Equal(t, expectedLalamoveLog.Fields, actualLalamoveLog.Fields)
	assert.Equal(t, expectedLalamoveLog.Level, actualLalamoveLog.Level)
}

// Copyright © 2017 Engine Yard, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/engineyard/eycore/debugging"
)

var lines []*Line

type Line struct {
	Time     time.Time
	Severity string
	Tag      string
	Message  string
}

func (line *Line) String() string {
	return fmt.Sprintf("%s - %s - [%s] - %s",
		line.Time.String(),
		strings.ToUpper(line.Severity),
		line.Tag,
		line.Message)
}

func init() {
	if lines == nil {
		lines = make([]*Line, 0)
	}
}

func Record(severity string, tag string, message string) {
	if lines == nil {
		lines = make([]*Line, 0)
	}

	line := &Line{Time: time.Now(), Severity: severity, Tag: tag, Message: message}
	lines = append(lines, line)
	if debugging.Live() {
		fmt.Println(line)
	}
}

func Info(tag string, message string) {
	Record("INFO", tag, message)
}

func Warn(tag string, message string) {
	Record("WARN", tag, message)
}

func Debug(tag string, message string) {
	Record("DEBUG", tag, message)
}

func Error(tag string, message string) {
	Record("ERROR", tag, message)
}

func Lines() []string {
	printable := make([]string, 0)

	for _, line := range lines {
		printable = append(printable, line.String())
	}

	return printable
}

package http

import (
	"fmt"

	"github.com/engineyard/scaley/v2/pkg/scaley"
)

// LogService is a service that provides a logging mechanism via the Engine
// Yard alerts API.
type LogService struct {
	reportingURL string
}

// NewLogService takes a reporting url and returns a new LogService configured
// to report to the Engine Yard Alerts API at that url.
func NewLogService(reportingURL string) *LogService {
	return &LogService{reportingURL}
}

// Info takes a group and a message and submits an informational alert to the
// associated alerts API.
func (service *LogService) Info(group *scaley.Group, message string) {
	notify(
		service.reportingURL,
		warning,
		normalize(group, message),
	)
}

// Success takes a group and a message and submits a success alert to the
// associated alerts API.
func (service *LogService) Success(group *scaley.Group, message string) {
	notify(
		service.reportingURL,
		okay,
		normalize(group, message),
	)
}

// Failure takes a group and a message and submits a failure alert to the
// associated alerts API.
func (service *LogService) Failure(group *scaley.Group, message string) {
	notify(
		service.reportingURL,
		failure,
		normalize(group, message),
	)
}

func normalize(group *scaley.Group, message string) string {
	return fmt.Sprintf("Group[%s]: %s", group.Name, message)
}

// Copyright © 2019 Engine Yard, Inc.
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

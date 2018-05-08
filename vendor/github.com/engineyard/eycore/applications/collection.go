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

package applications

// Collection is a wrapper that helps decode API results that return
// several Server records.
type Collection struct {
	//Servers   []*Server `json:"servers"`
	Collected []*Model `json:"applications"`

	// Fields not marshalled to/from JSON
	AccountID     int `json:"-"`
	EnvironmentID int `json:"-"`
}
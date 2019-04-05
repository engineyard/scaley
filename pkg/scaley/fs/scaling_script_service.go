package fs

// ScalingScriptService is a service that knows how to interact with scaling
// scripts via the file system.
type ScalingScriptService struct{}

// NewScalingScriptService returns a new ScalingScriptService.
func NewScalingScriptService() *ScalingScriptService {
	return &ScalingScriptService{}
}

// Exists takes a path and returns a boolean that expresses whether or not the
// scaling script at that location exists.
func (service *ScalingScriptService) Exists(path string) bool {
	return FileExists(path)
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

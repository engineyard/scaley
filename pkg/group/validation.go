// Copyright © 2018 Engine Yard, Inc.
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

package group

import (
	"fmt"

	"github.com/engineyard/scaley/pkg/common"
)

func Validate(group *Group) error {

	if len(group.ScalingScript) == 0 {
		return fmt.Errorf("%s requires a scaling_script", group.Name)
	}

	if !common.FileExists(group.ScalingScript) {
		return fmt.Errorf("%s scaling_script does not exist on the file system", group.Name)
	}

	if len(group.ScalingServers) < 1 {
		return fmt.Errorf("%s contains no scaling servers", group.Name)
	}

	for _, server := range group.ScalingServers {
		if server.Instance == nil {
			return fmt.Errorf("%s contains invalid scaling server %s", group.Name, server.ID)
		}
	}

	return nil
}

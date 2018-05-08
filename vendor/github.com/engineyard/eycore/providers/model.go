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

package providers

import "fmt"

type credentials struct {
	InstanceAwsSecretID  string `json:"instance_aws_secret_id,omitempty"`
	InstanceAwsSecretKey string `json:"instance_aws_secret_key,omitempty"`
	AwsSecretID          string `json:"aws_secret_id"`
	AwsSecretKey         string `json:"aws_secret_key,omitempty"`
	AwsLogin             string `json:"aws_login,omitempty"`
	AwsPass              string `json:"aws_pass,omitempty"`
	PayerAccountID       string `json:"payer_account_id,omitempty"`
}

// Model is an object that represents a cloud provider in the Engine Yard
// API.type Model struct {
type Model struct {
	ID            int          `json:"id,omitempty"`
	ProvisionedID string       `json:"provisioned_id,omitempty"`
	Type          string       `json:"type,omitempty"`
	Credentials   *credentials `json:"credentials,omitempty"`

	// Timestamps
	CancelledAt string `json:"cancelled_at,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`

	// Account
	AccountID string `json:"account_id,omitempty"`
}

func (model *Model) String() string {
	return fmt.Sprintf("%d, (Account: %s)", model.ID, model.AccountID)
}

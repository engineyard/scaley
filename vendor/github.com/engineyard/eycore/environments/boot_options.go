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

package environments

import (
	"encoding/json"
	"strconv"
)

type appsConfig struct {
	Count         int    `json:"count,omitempty"`
	Flavor        string `json:"flavor,omitempty"`
	VolumeSize    string `json:"volume_size,omitempty"`
	MntVolumeSize string `json:"mnt_volume_size,omitempty"`
}

type dbMasterConfig struct {
	Flavor        string `json:"flavor,omitempty"`
	VolumeSize    string `json:"volume_size,omitempty"`
	MntVolumeSize string `json:"mnt_volume_size,omitempty"`
	Iops          string `json:"iops,omitempty"`
}
type dbSlaveConfig struct {
	Name   string `json:"name,omitempty"`
	Flavor string `json:"flavor,omitempty"`
}

type utilConfig struct {
	Name          string `json:"name,omitempty"`
	Flavor        string `json:"flavor,omitempty"`
	VolumeSize    string `json:"volume_size,omitempty"`
	MntVolumeSize string `json:"mnt_volume_size,omitempty"`
	Iops          string `json:"iops,omitempty"`
}

type BootOptions struct {
	Type         string           `json:"type,omitempty"`
	InstanceSize string           `json:"instance_size,omitempty"`
	IPID         int              `json:"ip_id,omitempty"`
	Apps         *appsConfig      `json:"apps,omitempty"`
	DbMaster     *dbMasterConfig  `json:"db_master,omitempty"`
	DbSlaves     []*dbSlaveConfig `json:"db_slaves,omitempty"`
	Utils        []*utilConfig    `json:"utils,omitempty"`
}

func NewBootOptions() *BootOptions {
	return &BootOptions{}
}

func (p *BootOptions) SetIPID(id int) {
	p.IPID = id
}

func (p *BootOptions) SetType(etype string) {
	p.Type = etype
}

func (p *BootOptions) SetInstanceSize(size string) {
	p.InstanceSize = size
}

func (p *BootOptions) ConfigureApps(count int, flavor string, volumeSize int, mntVolumeSize int) {
	p.Apps = &appsConfig{
		Count:         count,
		Flavor:        flavor,
		VolumeSize:    strconv.Itoa(volumeSize),
		MntVolumeSize: strconv.Itoa(mntVolumeSize),
	}
}

func (p *BootOptions) ConfigureDbMaster(flavor string, volumeSize int, mntVolumeSize int, iops int) {
	p.DbMaster = &dbMasterConfig{
		Flavor:        flavor,
		VolumeSize:    strconv.Itoa(volumeSize),
		MntVolumeSize: strconv.Itoa(mntVolumeSize),
		Iops:          strconv.Itoa(iops),
	}
}

func (p *BootOptions) AddDbSlave(name string, flavor string) {
	newSlave := &dbSlaveConfig{
		Name:   name,
		Flavor: flavor,
	}

	p.DbSlaves = append(p.DbSlaves, newSlave)
}

func (p *BootOptions) AddUtil(name string, flavor string, volumeSize int, mntVolumeSize int, iops int) {
	newUtil := &utilConfig{
		Name:          name,
		Flavor:        flavor,
		VolumeSize:    strconv.Itoa(volumeSize),
		MntVolumeSize: strconv.Itoa(mntVolumeSize),
		Iops:          strconv.Itoa(iops),
	}

	p.Utils = append(p.Utils, newUtil)
}

type clusterConfigWrapper struct {
	Config *BootOptions `json:"configuration,omitempty"`
}

func (p *BootOptions) Body() []byte {
	var data []byte

	wrapper := struct {
		ClusterConfig *clusterConfigWrapper `json:"cluster_configuration,omitempty"`
	}{&clusterConfigWrapper{p}}

	data, _ = json.Marshal(wrapper)

	return data
}

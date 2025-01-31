/*
 Copyright 2021 The KubeSphere Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package config

import (
	"github.com/kubesphere/kubekey/pkg/common"
	"github.com/kubesphere/kubekey/pkg/core/task"
)

type ModifyConfigModule struct {
	common.KubeModule
}

func (m *ModifyConfigModule) Init() {
	m.Name = "ModifyConfigModule"
	m.Desc = "Modify the KubeKey config file"

	modify := &task.LocalTask{
		Name:   "ModifyConfig",
		Desc:   "Modify the KubeKey config file",
		Action: new(ModifyConfig),
	}

	m.Tasks = []task.Interface{
		modify,
	}
}

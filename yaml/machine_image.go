// Code generated by scripts/generate.js; DO NOT EDIT.

// Copyright 2022 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package yaml

import (
	"encoding/json"
	"fmt"
)

// MachineImage defines possible machine image values.
type MachineImage string

// MachineImage enumeration.
const (
	MachineImageNone         MachineImage = ""
	MachineImageUbuntuLatest MachineImage = "ubuntu-latest"
	MachineImageMacosLatest  MachineImage = "macos-latest"
	MachineImageWndowsLatest MachineImage = "wndows-latest"
)

// UnmarshalJSON unmashals a quoted json string to the enum value.
func (e *MachineImage) UnmarshalJSON(b []byte) error {
	var v string
	json.Unmarshal(b, &v)
	switch v {
	case "":
		*e = MachineImageNone
	case "ubuntu-latest":
		*e = MachineImageUbuntuLatest
	case "macos-latest":
		*e = MachineImageMacosLatest
	case "wndows-latest":
		*e = MachineImageWndowsLatest
	default:
		if IsExpression(v) {
			*e = MachineImage(v)
		} else {
			return fmt.Errorf("invalid MachineImage: %s", v)
		}
	}
	return nil
}
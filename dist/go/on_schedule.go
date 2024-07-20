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
	"errors"
)

type Schedule struct {
	Items []*ScheduleItem
}

type ScheduleItem struct {
	Cron string `yaml:"cron,omitempty"`
}

// UnmarshalJSON implement the json.Unmarshaler interface.
func (v *Schedule) UnmarshalJSON(data []byte) error {
	out1 := new(ScheduleItem)
	out2 := []*ScheduleItem{}

	if err := json.Unmarshal(data, &out1); err == nil {
		v.Items = append(v.Items, out1)
		return nil
	}
	if err := json.Unmarshal(data, &out2); err == nil {
		v.Items = append(v.Items, out2...)
		return nil
	}

	return errors.New("failed to unmarshal on.schedule")
}
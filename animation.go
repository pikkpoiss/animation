// Copyright 2016 Pikkpoiss Authors
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

package animation

import (
	"time"
)

const (
	Step60Hz = time.Duration(16666) * time.Microsecond
	Step30Hz = Step60Hz * 2
	Step20Hz = time.Duration(50000) * time.Microsecond
	Step15Hz = Step30Hz * 2
	Step10Hz = Step20Hz * 2
	Step5Hz  = Step10Hz * 2
)

type AnimationCallback func()

type Animation struct {
	elapsed  time.Duration
	callback AnimationCallback
}

func NewAnimation() *Animation {
	return &Animation{}
}

func (a *Animation) Update(elapsed time.Duration) (done bool) {
	a.elapsed += elapsed
	return false
}

func (a *Animation) Elapsed() time.Duration {
	return a.elapsed
}

func (a *Animation) Reset() {
	a.elapsed = time.Duration(0)
}

func (a *Animation) SetCallback(callback AnimationCallback) {
	if a.callback != nil {
		a.callback()
	}
	a.callback = callback
}

func (a *Animation) HasCallback() bool {
	return a.callback != nil
}

func (a *Animation) Callback() {
	if a.HasCallback() {
		a.callback()
	}
}

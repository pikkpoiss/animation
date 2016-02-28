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

type BoundedAnimation struct {
	Elapsed  time.Duration
	Duration time.Duration
	callback AnimatorCallback
}

func NewBoundedAnimation(duration time.Duration) *BoundedAnimation {
	return &BoundedAnimation{0, duration, nil}
}

func (a *BoundedAnimation) SetCallback(callback AnimatorCallback) {
	a.callback = callback
}

func (a *BoundedAnimation) IsDone() bool {
	return a.Elapsed >= a.Duration
}

func (a *BoundedAnimation) Update(elapsed time.Duration) time.Duration {
	a.Elapsed += elapsed
	if a.IsDone() {
		if a.callback != nil {
			a.callback()
		}
		return a.Elapsed - a.Duration
	}
	return 0
}

func (a *BoundedAnimation) Reset() {
	a.Elapsed = 0
}

func (a *BoundedAnimation) Delete() {
}

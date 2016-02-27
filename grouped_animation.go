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

type GroupedAnimation struct {
	animators []Animator
	callback  AnimatorCallback
}

func NewGroupedAnimation(animators []Animator) *GroupedAnimation {
	return &GroupedAnimation{animators, nil}
}

func (a *GroupedAnimation) SetCallback(callback AnimatorCallback) {
	a.callback = callback
}

func (a *GroupedAnimation) IsDone() bool {
	var done = true
	for _, animator := range a.animators {
		if !animator.IsDone() {
			done = false
		}
	}
	return done
}

func (a *GroupedAnimation) Update(elapsed time.Duration) time.Duration {
	var (
		total     time.Duration
		remainder time.Duration
		done      = true
	)
	for _, animator := range a.animators {
		remainder = animator.Update(elapsed)
		if !animator.IsDone() {
			done = false
		}
		if remainder != 0 && (total == 0 || remainder < total) {
			total = remainder // Take the smallest nonzero remainder.
		}
	}
	if done {
		if a.callback != nil {
			a.callback()
		}
		return total
	}
	return 0
}

func (a *GroupedAnimation) Reset() {
	for _, animator := range a.animators {
		animator.Reset()
	}
}

func (a *GroupedAnimation) Delete() {
	for _, animator := range a.animators {
		animator.Delete()
	}
	a.animators = []Animator{}
}

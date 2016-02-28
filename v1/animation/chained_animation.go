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

type ChainedAnimation struct {
	animators []Animator
	loop      bool
	index     int
	callback  AnimatorCallback
}

func NewChainedAnimation(animators []Animator, loop bool) *ChainedAnimation {
	return &ChainedAnimation{animators, loop, 0, nil}
}

func (a *ChainedAnimation) SetCallback(callback AnimatorCallback) {
	a.callback = callback
}

func (a *ChainedAnimation) IsDone() bool {
	var count = len(a.animators)
	return !a.loop && count > 0 && a.animators[count-1].IsDone()
}

func (a *ChainedAnimation) Update(elapsed time.Duration) time.Duration {
	var count = len(a.animators)
	if count > a.index {
		for elapsed > 0 && !a.animators[a.index].IsDone() {
			elapsed = a.animators[a.index].Update(elapsed)
			if a.animators[a.index].IsDone() {
				if a.loop {
					a.animators[a.index].Reset()
				}
				a.index = (a.index + 1) % count
				if !a.loop && a.index == 0 && a.callback != nil {
					a.callback()
					break
				}
			}
		}
	}
	return elapsed
}

func (a *ChainedAnimation) Reset() {
	a.index = 0
	for _, animator := range a.animators {
		animator.Reset()
	}
}

func (a *ChainedAnimation) Delete() {
	for _, animator := range a.animators {
		animator.Delete()
	}
	a.animators = []Animator{}
}

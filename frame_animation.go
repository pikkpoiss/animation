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

type FrameAnimation struct {
	*Animation
	FrameLength time.Duration
	Sequence    []int
	Current     int
}

func NewFrameAnimation(length time.Duration, frames []int) *FrameAnimation {
	return &FrameAnimation{
		Animation:   NewAnimation(),
		FrameLength: length,
		Sequence:    frames,
		Current:     frames[0],
	}
}

func (a *FrameAnimation) Update(elapsed time.Duration) (done bool) {
	a.Animation.Update(elapsed)
	index := int(a.Elapsed()/a.FrameLength) % len(a.Sequence)
	a.Current = a.Sequence[index]
	done = false
	if a.HasCallback() && index == len(a.Sequence)-1 {
		a.Callback()
		a.SetCallback(nil)
		done = true
	}
	return
}

func (a *FrameAnimation) OffsetFrame(offset int) int {
	index := int(a.Elapsed()/a.FrameLength) % len(a.Sequence)
	return a.Sequence[(index+offset)%len(a.Sequence)]
}

func (a *FrameAnimation) SetSequence(seq []int) {
	a.Sequence = seq
	a.Current = a.Sequence[0]
	a.Animation.Reset()
}

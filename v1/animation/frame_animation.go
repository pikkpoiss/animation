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

type Frame struct {
	Duration time.Duration
	Index    int
}

func MsFrame(milliseconds, index int) Frame {
	return Frame{time.Duration(milliseconds) * time.Millisecond, index}
}

type FrameAnimation struct {
	Elapsed   time.Duration
	remainder time.Duration
	Duration  time.Duration
	callback  AnimatorCallback
	sequence  []Frame
	current   int
	loop      bool
	target    *int
}

func NewFrameAnimation(frames []Frame, loop bool, target *int) *FrameAnimation {
	var f = &FrameAnimation{
		loop:   loop,
		target: target,
	}
	f.SetFrames(frames)
	return f
}

func (a *FrameAnimation) IsDone() bool {
	return !a.loop && a.Elapsed >= a.Duration
}

func (a *FrameAnimation) SetCallback(callback AnimatorCallback) {
	a.callback = callback
}

func (a *FrameAnimation) Update(elapsed time.Duration) time.Duration {
	a.Elapsed += elapsed
	elapsed += a.remainder
	for elapsed > 0 {
		if elapsed >= a.sequence[a.current].Duration {
			elapsed -= a.sequence[a.current].Duration
			a.current += 1
			if a.current >= len(a.sequence) {
				if a.loop {
					a.current = a.current % len(a.sequence)
				} else {
					a.current = len(a.sequence) - 1
				}
			}
		} else {
			a.remainder = elapsed
			elapsed = 0
		}
	}
	if a.target != nil {
		*a.target = a.sequence[a.current].Index
	}
	if a.IsDone() {
		if a.callback != nil {
			a.callback()
		}
		return a.Elapsed - a.Duration
	}
	return 0
}

func (a *FrameAnimation) Reset() {
	a.current = 0
	a.Elapsed = 0
}

func (a *FrameAnimation) Delete() {}

func (a *FrameAnimation) SetFrames(frames []Frame) {
	var (
		duration time.Duration = 0
	)
	for _, frame := range frames {
		duration += frame.Duration
	}
	a.Duration = duration
	a.sequence = frames
	a.Reset()
}

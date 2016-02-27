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
	"math"
	"time"
)

// Returns {value}, {done}, {remainder}
type ContinuousFunc func(elapsed time.Duration) (float32, bool, time.Duration)

type ContinuousAnimation struct {
	Elapsed  time.Duration
	function ContinuousFunc
	target   *float32
	callback AnimatorCallback
	done     bool
}

func NewContinuousAnimation(f ContinuousFunc, target *float32) *ContinuousAnimation {
	return &ContinuousAnimation{
		function: f,
		target:   target,
		done:     false,
	}
}

func (a *ContinuousAnimation) Update(elapsed time.Duration) time.Duration {
	var (
		remainder time.Duration
		result    float32
	)
	a.Elapsed += elapsed
	result, a.done, remainder = a.function(a.Elapsed)
	if a.target != nil {
		*a.target = result
	}
	if a.IsDone() {
		if a.callback != nil {
			a.callback()
		}
	}
	return remainder
}

func (a *ContinuousAnimation) SetCallback(callback AnimatorCallback) {
	a.callback = callback
}

func (a *ContinuousAnimation) IsDone() bool {
	return a.done
}

func (a *ContinuousAnimation) Reset() {
	a.done = false
	a.Elapsed = 0
}

func (a *ContinuousAnimation) Delete() {}

func SineDecayFunc(duration time.Duration, amplitude, frequency, decay float32) ContinuousFunc {
	var interval = float64(frequency * 2.0 * math.Pi)
	return func(elapsed time.Duration) (value float32, done bool, remainder time.Duration) {
		done = elapsed >= duration
		remainder = elapsed - duration
		if done {
			value = 0
		} else {
			decayAmount := 1.0 - float32(elapsed)/float32(duration)*decay
			value = float32(math.Sin(elapsed.Seconds()*interval/duration.Seconds())) * amplitude * decayAmount
		}
		return
	}
}

func LinearFunc(duration time.Duration, from, to float32) ContinuousFunc {
	return func(elapsed time.Duration) (value float32, done bool, remainder time.Duration) {
		var (
			denom = float64(duration)
			numer = math.Min(float64(elapsed), denom)
			pct   = float32(numer / denom)
		)
		value = pct*(to-from) + from
		done = elapsed >= duration
		remainder = elapsed - duration
		return
	}
}

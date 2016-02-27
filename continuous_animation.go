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

type ContinuousFunc func(elapsed time.Duration) float32

type ContinuousAnimation struct {
	*Animation
	function ContinuousFunc
	target   *float32
}

func NewContinuousAnimation(f ContinuousFunc, target *float32) *ContinuousAnimation {
	return &ContinuousAnimation{
		Animation: NewAnimation(),
		function:  f,
		target:    target,
	}
}

func (a *ContinuousAnimation) Update(elapsed time.Duration) time.Duration {
	var (
		//leftover = a.Animation.Update(elapsed)
	)
	a.Animation.Update(elapsed)
	*a.target = a.function(a.Elapsed())
	//return leftover
	return 1 * time.Second
}

func SineDecayFunc(duration time.Duration, amplitude, frequency, decay float32, callback AnimationCallback) ContinuousFunc {
	var interval = float64(frequency * 2.0 * math.Pi)
	return func(elapsed time.Duration) float32 {
		if elapsed > duration {
			if callback != nil {
				callback()
			}
			return 0.0
		}
		decayAmount := 1.0 - float32(elapsed)/float32(duration)*decay
		return float32(math.Sin(elapsed.Seconds()*interval/duration.Seconds())) * amplitude * decayAmount
	}
}

func LinearFunc(from, to float32, duration time.Duration) ContinuousFunc {
	return func(elapsed time.Duration) float32 {
		var (
			denom = float64(duration)
			numer = math.Min(float64(elapsed), denom)
			pct   = float32(numer / denom)
		)
		return pct*(to-from) + from
	}
}

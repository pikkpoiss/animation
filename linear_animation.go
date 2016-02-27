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

type LinearAnimation struct {
	BoundedAnimation
	target *float32
	from   float32
	to     float32
}

func NewLinearAnimation(target *float32, from, to float32, duration time.Duration) *LinearAnimation {
	return &LinearAnimation{
		BoundedAnimation{
			0,
			duration,
			nil,
		},
		target,
		from,
		to,
	}
}

func (a *LinearAnimation) Update(elapsed time.Duration) {
	a.BoundedAnimation.Update(elapsed)
	var (
		denom = float64(a.BoundedAnimation.Duration)
		numer = math.Min(float64(a.BoundedAnimation.Elapsed), denom)
		pct   = float32(numer / denom)
		value = pct*(a.to-a.from) + a.from
	)
	*a.target = value
}

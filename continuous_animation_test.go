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
	"testing"
	"time"
)

// Tests that a continuous animation with sine decay function produces expected values.
func TestContinuousAnimationSineDecay(t *testing.T) {
	var (
		anim = NewContinuousAnimation(SineDecayFunc(3 * time.Second, 5, 1, 1, nil))
	)
	anim.Update(1 * time.Second)
	if anim.Value() != 2.886751 {
		t.Errorf("SineDecayFunc produced unexpected value %v", anim.Value())
	}
	anim.Update(1 * time.Second)
	if anim.Value() != -1.4433757 {
		t.Errorf("SineDecayFunc produced unexpected value %v", anim.Value())
	}
	anim.Update(1 * time.Second)
	if anim.Value() != 0 {
		t.Errorf("SineDecayFunc produced unexpected value %v", anim.Value())
	}
	anim.Update(1 * time.Second)
	if anim.Value() != 0 {
		t.Errorf("SineDecayFunc produced unexpected value %v", anim.Value())
	}
}

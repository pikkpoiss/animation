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
		dest float32 = 0.0
		anim         = NewContinuousAnimation(SineDecayFunc(3*time.Second, 5, 1, 1), &dest)
	)
	anim.Update(1 * time.Second)
	if dest != 2.886751 {
		t.Errorf("SineDecayFunc produced unexpected value %v", dest)
	}
	anim.Update(1 * time.Second)
	if dest != -1.4433757 {
		t.Errorf("SineDecayFunc produced unexpected value %v", dest)
	}
	anim.Update(1 * time.Second)
	if dest != 0 {
		t.Errorf("SineDecayFunc produced unexpected value %v", dest)
	}
	anim.Update(1 * time.Second)
	if dest != 0 {
		t.Errorf("SineDecayFunc produced unexpected value %v", dest)
	}
}

// Tests that a linear animation produces expected values.
func TestContinuousLinearAnimation(t *testing.T) {
	var (
		dest float32 = 1.0
		anim         = NewContinuousAnimation(LinearFunc(5*time.Second, 10, 20), &dest)
	)
	if dest != 1.0 {
		t.Fatalf("Target value does not match expected, got %v", dest)
	}
	anim.Update(1000 * time.Millisecond)
	if dest != 12 {
		t.Fatalf("Target value does not match expected, got %v", dest)
	}
	anim.Update(1000 * time.Millisecond)
	if dest != 14 {
		t.Fatalf("Target value does not match expected, got %v", dest)
	}
	anim.Update(1000 * time.Millisecond)
	if dest != 16 {
		t.Fatalf("Target value does not match expected, got %v", dest)
	}
	anim.Update(1000 * time.Millisecond)
	if dest != 18 {
		t.Fatalf("Target value does not match expected, got %v", dest)
	}
	anim.Update(1000 * time.Millisecond)
	if dest != 20 {
		t.Fatalf("Target value does not match expected, got %v", dest)
	}
	anim.Update(1000 * time.Millisecond)
	if dest != 20 {
		t.Fatalf("Target value does not match expected, got %v", dest)
	}
}

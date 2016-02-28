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

// Tests that setting a callback works and the callback is called.
func TestBoundedAnimationSetCallback(t *testing.T) {
	var (
		done = false
		anim = BoundedAnimation{0, 1 * time.Second, nil}
		cb   = func() { done = true }
	)
	anim.SetCallback(cb)
	anim.Update(1 * time.Second)
	if done != true {
		t.Fatalf("BoundedAnimation callback not called")
	}
}

// Tests that BoundedAnimation.IsDone is true when elapsed == duration.
func TestBoundedAnimationIsDoneElapsedEqDuration(t *testing.T) {
	var anim = BoundedAnimation{0, 1 * time.Second, nil}
	anim.Update(1 * time.Second)
	if !anim.IsDone() {
		t.Fatalf("BoundedAnimation.IsDone was not expected value")
	}
}

// Tests that BoundedAnimation.IsDone is true when elapsed > duration.
func TestBoundedAnimationIsDoneElapsedGtDuration(t *testing.T) {
	var anim = BoundedAnimation{0, 1 * time.Second, nil}
	anim.Update(9 * time.Second)
	if !anim.IsDone() {
		t.Fatalf("BoundedAnimation.IsDone was not expected value")
	}
}

// Tests that BoundedAnimation.IsDone is false when elapsed < duration.
func TestBoundedAnimationIsDoneElapsedLtDuration(t *testing.T) {
	var anim = BoundedAnimation{0, 1 * time.Second, nil}
	anim.Update(200 * time.Millisecond)
	if anim.IsDone() {
		t.Fatalf("BoundedAnimation.IsDone was not expected value")
	}
}

// Tests that BoundedAnimation.Update increments elapsed and returns overflow.
func TestBoundedAnimationUpdate(t *testing.T) {
	var (
		anim               = BoundedAnimation{0, 1 * time.Second, nil}
		resp time.Duration = 0
	)
	resp = anim.Update(200 * time.Millisecond)
	if anim.Elapsed != 200*time.Millisecond {
		t.Fatalf("BoundedAnimation.Elapsed was not expected value")
	}
	if resp != 0 {
		t.Fatalf("BoundedAnimation.Elapsed did not return expected value")
	}
	resp = anim.Update(600 * time.Millisecond)
	if anim.Elapsed != 800*time.Millisecond {
		t.Fatalf("BoundedAnimation.Elapsed was not expected value")
	}
	if resp != 0 {
		t.Fatalf("BoundedAnimation.Elapsed did not return expected value")
	}
	resp = anim.Update(400 * time.Millisecond)
	if anim.Elapsed != 1200*time.Millisecond {
		t.Fatalf("BoundedAnimation.Elapsed was not expected value")
	}
	if resp != 200*time.Millisecond {
		t.Fatalf("BoundedAnimation.Elapsed did not return expected value")
	}
}

func TestBoundedAnimationReset(t *testing.T) {
	var anim = BoundedAnimation{200 * time.Millisecond, 1 * time.Second, nil}
	anim.Reset()
	if anim.Elapsed != 0 {
		t.Fatalf("BoundedAnimation.Elapsed was not expected value")
	}
}

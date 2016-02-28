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

func TestLoopFrameAnimation(t *testing.T) {
	var (
		target int  = 0
		done   bool = false
		frames      = []Frame{MsFrame(100, 0), MsFrame(100, 2), MsFrame(100, 1), MsFrame(100, 3)}
		anim        = NewFrameAnimation(frames, true, &target)
		cb          = func() { done = true }
	)
	anim.SetCallback(cb)
	anim.Update(50 * time.Millisecond)
	if target != 0 {
		t.Fatalf("Current frame does not match expected, got %v", target)
	}
	anim.Update(100 * time.Millisecond)
	if target != 2 {
		t.Fatalf("Current frame does not match expected, got %v", target)
	}
	anim.Update(100 * time.Millisecond)
	if target != 1 {
		t.Fatalf("Current frame does not match expected, got %v", target)
	}
	anim.Update(100 * time.Millisecond)
	if target != 3 {
		t.Fatalf("Current frame does not match expected, got %v", target)
	}
	anim.Update(100 * time.Millisecond)
	if target != 0 {
		t.Fatalf("Current frame does not match expected, got %v", target)
	}
	anim.Update(100 * time.Millisecond)
	if target != 2 {
		t.Fatalf("Current frame does not match expected, got %v", target)
	}
	if anim.IsDone() {
		t.Fatalf("Looping animation marked done (should not)")
	}
	if done {
		t.Fatalf("Non-looping animation called callback (should not)")
	}
}

func TestNonLoopFrameAnimation(t *testing.T) {
	var (
		target int  = 0
		done   bool = false
		frames      = []Frame{MsFrame(100, 0), MsFrame(100, 2), MsFrame(100, 1), MsFrame(100, 3)}
		anim        = NewFrameAnimation(frames, false, &target)
		cb          = func() { done = true }
	)
	anim.SetCallback(cb)
	anim.Update(50 * time.Millisecond)
	if target != 0 {
		t.Fatalf("Current frame does not match expected, got %v", target)
	}
	anim.Update(100 * time.Millisecond)
	if target != 2 {
		t.Fatalf("Current frame does not match expected, got %v", target)
	}
	anim.Update(100 * time.Millisecond)
	if target != 1 {
		t.Fatalf("Current frame does not match expected, got %v", target)
	}
	anim.Update(100 * time.Millisecond)
	if target != 3 {
		t.Fatalf("Current frame does not match expected, got %v", target)
	}
	if anim.IsDone() {
		t.Fatalf("Non-looping animation marked done too early")
	}
	if done {
		t.Fatalf("Non-looping animation called callback too early")
	}
	anim.Update(100 * time.Millisecond)
	if target != 3 {
		t.Fatalf("Current frame does not match expected, got %v", target)
	}
	if !anim.IsDone() {
		t.Fatalf("Non-looping animation not marked done when finished")
	}
	if !done {
		t.Fatalf("Non-looping animation did not call callback when done")
	}
	anim.Update(100 * time.Millisecond)
	if target != 3 {
		t.Fatalf("Current frame does not match expected, got %v", target)
	}
}

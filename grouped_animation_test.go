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

func TestGroupedAnimationSetCallback(t *testing.T) {
	var (
		done   = false
		child1 = &BoundedAnimation{0, 1 * time.Second, nil}
		child2 = &BoundedAnimation{0, 2 * time.Second, nil}
		anim   = GroupedAnimation{[]Animator{child1, child2}, nil}
		cb     = func() { done = true }
	)
	anim.SetCallback(cb)
	anim.Update(1 * time.Second)
	if done {
		t.Fatalf("GroupedAnimation callback called too early")
	}
	if !child1.IsDone() {
		t.Fatalf("First child animation not marked done")
	}
	if child2.IsDone() {
		t.Fatalf("Second child animation marked done too early")
	}
	anim.Update(1 * time.Second)
	if !done {
		t.Fatalf("GroupedAnimation callback not called")
	}
	if !child2.IsDone() {
		t.Fatalf("Second child animation not marked done")
	}
	if child1.Elapsed != 2*time.Second {
		t.Fatalf("First child elapsed not equal to expected value")
	}
	if child2.Elapsed != 2*time.Second {
		t.Fatalf("Second child elapsed not equal to expected value")
	}
}

func TestGroupedAnimationIsDone(t *testing.T) {
	var (
		child1 = &BoundedAnimation{0, 1 * time.Second, nil}
		child2 = &BoundedAnimation{0, 2 * time.Second, nil}
		anim   = GroupedAnimation{[]Animator{child1, child2}, nil}
	)
	anim.Update(1 * time.Second)
	if anim.IsDone() {
		t.Fatalf("GroupedAnimation.IsDone is true too early")
	}
	anim.Update(1 * time.Second)
	if !anim.IsDone() {
		t.Fatalf("GroupedAnimation.IsDone is not true after animation")
	}
}

func TestGroupedAnimationUpdate(t *testing.T) {
	var (
		child1               = &BoundedAnimation{0, 1 * time.Second, nil}
		child2               = &BoundedAnimation{0, 2 * time.Second, nil}
		anim                 = GroupedAnimation{[]Animator{child1, child2}, nil}
		resp   time.Duration = 0
	)
	resp = anim.Update(1 * time.Second)
	if resp != 0 {
		t.Fatalf("GroupedAnimation.Update did not return correct remainder")
	}
	resp = anim.Update(200 * time.Millisecond)
	if resp != 0 {
		t.Fatalf("GroupedAnimation.Update did not return correct remainder")
	}
	resp = anim.Update(900 * time.Millisecond)
	if resp != 100*time.Millisecond {
		t.Fatalf("GroupedAnimation.Update did not return correct remainder, got %v", resp)
	}
}

func TestGroupedAnimationReset(t *testing.T) {
	var (
		child1 = &BoundedAnimation{200 * time.Millisecond, 1 * time.Second, nil}
		child2 = &BoundedAnimation{1200 * time.Millisecond, 2 * time.Second, nil}
		anim   = GroupedAnimation{[]Animator{child1, child2}, nil}
	)
	anim.Reset()
	if child1.Elapsed != 0 {
		t.Fatalf("GroupedAnimation.Reset did not reset first child")
	}
	if child2.Elapsed != 0 {
		t.Fatalf("GroupedAnimation.Reset did not reset second child")
	}
}

func TestGroupedAnimationDelete(t *testing.T) {
	var (
		child1 = &BoundedAnimation{0, 1 * time.Second, nil}
		child2 = &BoundedAnimation{0, 2 * time.Second, nil}
		anim   = GroupedAnimation{[]Animator{child1, child2}, nil}
	)
	anim.Delete()
	anim.Update(100 * time.Millisecond)
	if child1.Elapsed != 0 {
		t.Fatalf("GroupedAnimation.Delete must remove references to children")
	}
}

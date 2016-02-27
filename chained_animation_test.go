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

func TestChainedAnimationSetCallback(t *testing.T) {
	var (
		done   = false
		child1 = &BoundedAnimation{0, 1 * time.Second, nil}
		child2 = &BoundedAnimation{0, 2 * time.Second, nil}
		anim   = ChainedAnimation{[]Animator{child1, child2}, false, 0, nil}
		cb     = func() { done = true }
	)
	anim.SetCallback(cb)
	anim.Update(1 * time.Second)
	if done {
		t.Fatalf("ChainedAnimation callback called too early")
	}
	anim.Update(2 * time.Second)
	if !done {
		t.Fatalf("ChainedAnimation callback not called")
	}
}

func TestChainedAnimationIsDoneNoLoop(t *testing.T) {
	var (
		child1 = &BoundedAnimation{0, 1 * time.Second, nil}
		child2 = &BoundedAnimation{0, 2 * time.Second, nil}
		anim   = ChainedAnimation{[]Animator{child1, child2}, false, 0, nil}
	)
	if anim.IsDone() {
		t.Fatalf("ChainedAnimation.IsDone true too early")
	}
	anim.Update(1 * time.Second)
	if anim.IsDone() {
		t.Fatalf("ChainedAnimation.IsDone true too early")
	}
	anim.Update(1 * time.Second)
	if anim.IsDone() {
		t.Fatalf("ChainedAnimation.IsDone true too early")
	}
	anim.Update(1 * time.Second)
	if !anim.IsDone() {
		t.Fatalf("ChainedAnimation.IsDone not set after animation finishes")
	}
}

func TestChainedAnimationIsDoneLoop(t *testing.T) {
	var (
		child1 = &BoundedAnimation{0, 1 * time.Second, nil}
		child2 = &BoundedAnimation{0, 2 * time.Second, nil}
		anim   = ChainedAnimation{[]Animator{child1, child2}, true, 0, nil}
	)
	if anim.IsDone() {
		t.Fatalf("ChainedAnimation.IsDone should not be true for loops")
	}
	anim.Update(1 * time.Second)
	if anim.IsDone() {
		t.Fatalf("ChainedAnimation.IsDone should not be true for loops")
	}
	anim.Update(1 * time.Second)
	if anim.IsDone() {
		t.Fatalf("ChainedAnimation.IsDone should not be true for loops")
	}
	anim.Update(1 * time.Second)
	if anim.IsDone() {
		t.Fatalf("ChainedAnimation.IsDone should not be true for loops")
	}
	anim.Update(1 * time.Second)
}

func TestChainedAnimationUpdateNoLoop(t *testing.T) {
	var (
		child1 = &BoundedAnimation{0, 1 * time.Second, nil}
		child2 = &BoundedAnimation{0, 2 * time.Second, nil}
		anim   = ChainedAnimation{[]Animator{child1, child2}, false, 0, nil}
		resp   time.Duration
	)
	resp = anim.Update(500 * time.Millisecond)
	if child1.Elapsed != 500*time.Millisecond {
		t.Fatalf("ChainedAnimation.Update must update current child")
	}
	if resp != 0 {
		t.Fatalf("ChainedAnimation.Update must return 0 before animation is done")
	}
	resp = anim.Update(600 * time.Millisecond)
	if child1.Elapsed != 1100*time.Millisecond {
		t.Fatalf("ChainedAnimation.Update must update current child")
	}
	if child2.Elapsed != 100*time.Millisecond {
		t.Fatalf("ChainedAnimation.Update must move to second child if first overflows")
	}
	if resp != 0 {
		t.Fatalf("ChainedAnimation.Update must return 0 before animation is done")
	}
	resp = anim.Update(500 * time.Millisecond)
	if child1.Elapsed != 1100*time.Millisecond {
		t.Fatalf("ChainedAnimation.Update must not update previous child")
	}
	if child2.Elapsed != 600*time.Millisecond {
		t.Fatalf("ChainedAnimation.Update must update current child")
	}
	if resp != 0 {
		t.Fatalf("ChainedAnimation.Update must return 0 before animation is done")
	}
	resp = anim.Update(1500 * time.Millisecond)
	if child1.Elapsed != 1100*time.Millisecond {
		t.Fatalf("ChainedAnimation.Update must update current child")
	}
	if child2.Elapsed != 2100*time.Millisecond {
		t.Fatalf("ChainedAnimation.Update must update current child")
	}
	if resp != 100*time.Millisecond {
		t.Fatalf("ChainedAnimation.Update must return remainder when animation is done")
	}
}

func TestChainedAnimationUpdateLoop(t *testing.T) {
	var (
		child1 = &BoundedAnimation{0, 1 * time.Second, nil}
		child2 = &BoundedAnimation{0, 2 * time.Second, nil}
		anim   = ChainedAnimation{[]Animator{child1, child2}, true, 0, nil}
		resp   time.Duration
	)
	resp = anim.Update(500 * time.Millisecond)
	if child1.Elapsed != 500*time.Millisecond {
		t.Fatalf("ChainedAnimation.Update must update current child")
	}
	if resp != 0 {
		t.Fatalf("ChainedAnimation.Update must return 0 for loops")
	}
	resp = anim.Update(600 * time.Millisecond)
	if child1.Elapsed != 0 {
		t.Fatalf("ChainedAnimation.Update must reset previous child")
	}
	if child2.Elapsed != 100*time.Millisecond {
		t.Fatalf("ChainedAnimation.Update must move to second child if first overflows")
	}
	if resp != 0 {
		t.Fatalf("ChainedAnimation.Update must return 0 for loops")
	}
	resp = anim.Update(500 * time.Millisecond)
	if child1.Elapsed != 0 {
		t.Fatalf("ChainedAnimation.Update must not update previous child")
	}
	if child2.Elapsed != 600*time.Millisecond {
		t.Fatalf("ChainedAnimation.Update must update current child")
	}
	if resp != 0 {
		t.Fatalf("ChainedAnimation.Update must return 0 for loops")
	}
	resp = anim.Update(1400 * time.Millisecond)
	if child1.Elapsed != 0 {
		t.Fatalf("ChainedAnimation.Update must not update previous child")
	}
	if child2.Elapsed != 0 {
		t.Fatalf("ChainedAnimation.Update must reset child when done")
	}
	if resp != 0 {
		t.Fatalf("ChainedAnimation.Update must return 0 for loops")
	}
	resp = anim.Update(300 * time.Millisecond)
	if child1.Elapsed != 300*time.Millisecond {
		t.Fatalf("ChainedAnimation.Update must loop back to first child")
	}
	if child2.Elapsed != 0 {
		t.Fatalf("ChainedAnimation.Update must update current child")
	}
	if resp != 0 {
		t.Fatalf("ChainedAnimation.Update must return 0 for loops")
	}
}

func TestChainedAnimationReset(t *testing.T) {
	var (
		child1 = &BoundedAnimation{100 * time.Millisecond, 1 * time.Second, nil}
		child2 = &BoundedAnimation{200 * time.Millisecond, 2 * time.Second, nil}
		anim   = ChainedAnimation{[]Animator{child1, child2}, true, 0, nil}
	)
	anim.Reset()
	if child1.Elapsed != 0 {
		t.Fatalf("ChainedAnimation.Reset must reset first child")
	}
	if child2.Elapsed != 0 {
		t.Fatalf("ChainedAnimation.Reset must reset second child")
	}
}

func TestChainedAnimationDelete(t *testing.T) {
	var (
		child1 = &BoundedAnimation{0, 1 * time.Second, nil}
		child2 = &BoundedAnimation{0, 2 * time.Second, nil}
		anim   = ChainedAnimation{[]Animator{child1, child2}, true, 0, nil}
	)
	anim.Delete()
	anim.Update(100 * time.Millisecond)
	if child1.Elapsed != 0 {
		t.Fatalf("ChainedAnimation.Delete must remove references to children")
	}
}

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

func TestFrameAnimation(t *testing.T) {
	var (
		anim = NewFrameAnimation(100*time.Millisecond, []int{0, 2, 1, 3})
	)
	anim.Update(50 * time.Millisecond)
	if anim.Current != 0 {
		t.Fatalf("Current frame does not match expected, got %v", anim.Current)
	}
	anim.Update(100 * time.Millisecond)
	if anim.Current != 2 {
		t.Fatalf("Current frame does not match expected, got %v", anim.Current)
	}
	anim.Update(100 * time.Millisecond)
	if anim.Current != 1 {
		t.Fatalf("Current frame does not match expected, got %v", anim.Current)
	}
	anim.Update(100 * time.Millisecond)
	if anim.Current != 3 {
		t.Fatalf("Current frame does not match expected, got %v", anim.Current)
	}
	anim.Update(100 * time.Millisecond)
	if anim.Current != 0 {
		t.Fatalf("Current frame does not match expected, got %v", anim.Current)
	}
}

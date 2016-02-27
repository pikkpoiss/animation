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

// Tests that all animation classes satisfy the Animator interface.
func TestInterfaces(t *testing.T) {
	var anim Animator
	anim = NewBoundedAnimation(1 * time.Second)
	anim = NewChainedAnimation([]Animator{}, false)
	anim = NewGroupedAnimation([]Animator{})
	anim = NewFrameAnimation([]Frame{MsFrame(100, 0), MsFrame(100, 1), MsFrame(100, 2)}, false, nil)
	t.Logf("Done checking interfaces for %v", anim)
}

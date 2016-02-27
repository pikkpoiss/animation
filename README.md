# Animation

Utilities for managing animations.

Originally part of https://github.com/pikkpoiss/twodee but
broken out since this code can be used in other places.

## Interface

Animations must satisfy the `Animator` interface:

```
type Animator interface {
	SetCallback(callback AnimatorCallback)
	IsDone() bool
	Update(elapsed time.Duration) time.Duration
	Reset()
	Delete()
}
```

Advance an animation by calling `Update` with the amount of elapsed time.

The behavior of an animation depends on its concrete type.  You can supply
a `func()` callback to `SetCallback` to be called when the animation completes.

## Types of animations

### BoundedAnimation

An animation which is bounded by a fixed amount of time.

```
var (
	done = false
	anim = NewBoundedAnimation(1 * time.Second)
	cb   = func() { done = true }
)
anim.SetCallback(cb)
anim.Update(1 * time.Second)
assert(done == true)
```

### ChainedAnimation

Connects two animations serially.  When one animation
completes, the leftover time accrues into the second.
The callback to `ChainedAnimation` is only called when
all animations are done.

`ChainedAnimation` supports a loop parameter, which
will just cause the animation to loop infinitely.

```
var (
	done   = false
	child1 = NewBoundedAnimation(1 * time.Second)
	child2 = NewBoundedAnimation(2 * time.Second)
	anim   = NewChainedAnimation([]Animator{child1, child2}, false)
	cb     = func() { done = true }
)
anim.SetCallback(cb)
anim.Update(1 * time.Second)
assert(!done)
anim.Update(2 * time.Second)
assert(done)
```

### ContinuousAnimation

Supports animating a continuous value over time, as defined
by a function parameter:

```
// Returns {value}, {done}, {remainder}
type ContinuousFunc func(elapsed time.Duration) (float32, bool, time.Duration)
```

#### LinearFunc

Produces a linear interpolation between two values.

```
var (
	dest float32 = 1.0
	anim         = NewContinuousAnimation(LinearFunc(5*time.Second, 10, 20), &dest)
)
assert(dest == 1.0)
anim.Update(1000 * time.Millisecond)
assert(dest == 12)
anim.Update(1000 * time.Millisecond)
assert(dest == 14)
anim.Update(1000 * time.Millisecond)
assert(dest == 16)
anim.Update(1000 * time.Millisecond)
assert(dest == 18)
anim.Update(1000 * time.Millisecond)
assert(dest == 20)
anim.Update(1000 * time.Millisecond)
assert(dest == 20)
```

#### SineDecayFunc

Produces a decaying sine wave, like could be used for a camera
shake effect.

```
var (
	dest float32 = 0.0
	anim         = NewContinuousAnimation(SineDecayFunc(3*time.Second, 5, 1, 1), &dest)
)
anim.Update(1 * time.Second)
assert(dest == 2.886751)
anim.Update(1 * time.Second)
assert(dest == -1.4433757)
anim.Update(1 * time.Second)
assert(dest == 0)
anim.Update(1 * time.Second)
assert(dest == 0)
```

### FrameAnimation

An animation which iterates over a discrete sequence of frames.

```
var (
	target int  = 0
	frames      = []Frame{MsFrame(100, 0), MsFrame(100, 2), MsFrame(100, 1), MsFrame(100, 3)}
	anim        = NewFrameAnimation(frames, true, &target)
)
anim.Update(50 * time.Millisecond)
assert(target == 0)
anim.Update(100 * time.Millisecond)
assert(target == 2)
anim.Update(100 * time.Millisecond)
assert(target == 1)
anim.Update(100 * time.Millisecond)
assert(target == 3)
anim.Update(100 * time.Millisecond)
assert(target == 0)
```

### GroupedAnimation

A set of animations running in parallel.  The callback to
`GroupedAnimation` is only called when all animations have
finished.

```
var (
	done   = false
	child1 = NewBoundedAnimation(1 * time.Second)
	child2 = NewBoundedAnimation(2 * time.Second)
	anim   = NewGroupedAnimation([]Animator{child1, child2})
	cb     = func() { done = true }
)
anim.SetCallback(cb)
anim.Update(1 * time.Second)
assert(!done)
assert(child1.IsDone())
assert(!child2.IsDone())
anim.Update(1 * time.Second)
assert(done)
assert(child2.IsDone())
```

## Development

Run tests:

```
go test
```

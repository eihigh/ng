# ng

Package `ng` provides generic numeric types and interfaces. These facilitate interoperability between data structures that are structurally compatible, regardless of their original definition.

## Example

The following demonstrates how functions can use these interfaces to work with different, but structurally similar, types (e.g., `image.Point` and a custom struct).

```go
package ng_test

import (
	"image"
	"github.com/eihigh/ng"
)

// CirclesHit demonstrates interoperability using ng.Vec2like for vector inputs
// and ng.Scalar for generic scalar type S (used for components and radii).
func CirclesHit[P ng.Vec2like[S], S ng.Scalar](a, b P, rA, rB S) bool {
	va := ng.Vec2[S](a)
	vb := ng.Vec2[S](b)
	dx := va.X - vb.X
	dy := va.Y - vb.Y
	dsq := dx*dx + dy*dy
	return dsq <= (rA+rB)*(rA+rB)
}

type myVec2f struct{ X, Y float32 }

func Example_interoperability() {
	// Example call with image.Point
	var p1, p2 image.Point
	_ = CirclesHit(p1, p2, 1, 2)

	// Example call with a custom type myVec2f
	var f1, f2 myVec2f
	_ = CirclesHit(f1, f2, 1.0, 2.0)
}
```

For API details, see [pkg.go.dev/github.com/eihigh/ng](https://pkg.go.dev/github.com/eihigh/ng).

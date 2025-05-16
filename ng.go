// Package ng provides generic numeric types and interfaces
// for interoperability between structurally compatible types.
package ng

type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Int interface {
	SignedInt | UnsignedInt
}

type Float interface {
	~float32 | ~float64
}

type Scalar interface {
	Int | Float
}

type Vec2[S Scalar] struct{ X, Y S }
type Vec3[S Scalar] struct{ X, Y, Z S }
type Vec4[S Scalar] struct{ X, Y, Z, W S }

type Vec2like[S Scalar] interface{ ~struct{ X, Y S } }
type Vec3like[S Scalar] interface{ ~struct{ X, Y, Z S } }
type Vec4like[S Scalar] interface{ ~struct{ X, Y, Z, W S } }

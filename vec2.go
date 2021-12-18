package area

import "github.com/go-gl/mathgl/mgl64"

type Vec2 struct {
	minX, maxX, minY, maxY float64
}

func (v Vec2) Max() mgl64.Vec2 { return mgl64.Vec2{v.maxX, v.maxY} }
func (v Vec2) Min() mgl64.Vec2 { return mgl64.Vec2{v.minX, v.minY} }

func NewVec2(b1, b2 mgl64.Vec2) Vec2 {
	return Vec2{
		minX: minBound(b1.X(), b2.X()),
		maxX: maxBound(b1.X(), b2.X()),

		minY: minBound(b1.Y(), b2.Y()),
		maxY: maxBound(b1.Y(), b2.Y()),
	}
}

func (v Vec2) Vec2Within(vec mgl64.Vec2) bool {
	return vec.X() > v.minX && vec.X() < v.maxX && vec.Y() > v.minY && vec.Y() < v.maxY
}
func (v Vec2) Vec3WithinOrEqualXZ(vec mgl64.Vec3) bool {
	return vec.X() >= v.minX && vec.X() <= v.maxX && vec.Z() >= v.minY && vec.Z() <= v.maxY
}
func (v Vec2) Vec2WithinOrEqual(vec mgl64.Vec2) bool {
	return vec.X() >= v.minX && vec.X() <= v.maxX && vec.Y() >= v.minY && vec.Y() <= v.maxY
}
func (v Vec2) Vec3WithinXZ(vec mgl64.Vec3) bool {
	return vec.X() > v.minX && vec.X() < v.maxX && vec.Z() > v.minY && vec.Z() < v.maxY
}

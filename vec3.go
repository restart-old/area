package area

import "github.com/go-gl/mathgl/mgl64"

type Vec3 struct {
	minX, maxX, minY, maxY, minZ, maxZ float64
}

func (v Vec3) Max() mgl64.Vec3 { return mgl64.Vec3{v.maxX, v.maxY, v.maxZ} }
func (v Vec3) Min() mgl64.Vec3 { return mgl64.Vec3{v.minX, v.minY, v.minZ} }

func NewVec3(b1, b2 mgl64.Vec3) Vec3 {
	return Vec3{
		minX: minBound(b1.X(), b2.X()),
		maxX: maxBound(b1.X(), b2.X()),

		minY: minBound(b1.Y(), b2.Y()),
		maxY: maxBound(b1.Y(), b2.Y()),

		minZ: minBound(b1.Z(), b2.Z()),
		maxZ: maxBound(b1.Z(), b2.Z()),
	}
}
func (v Vec3) Vec3WithinOrEqual(vec mgl64.Vec3) bool {
	return vec.X() >= v.minX && vec.X() <= v.maxX && vec.Y() >= v.minY && vec.Y() <= v.maxY && vec.Z() >= v.minZ && vec.Z() <= v.maxZ
}

func (v Vec3) Vec3Within(vec mgl64.Vec3) bool {
	return vec.X() > v.minX && vec.X() < v.maxX && vec.Y() > v.minY && vec.Y() < v.maxY && vec.Z() > v.minZ && vec.Z() < v.maxZ
}

package area

import "github.com/go-gl/mathgl/mgl64"

type Area struct {
	minX, maxX, minY, maxY float64
}

func maxBound(b1, b2 float64) float64 {
	if b1 > b2 {
		return b1
	}
	return b2
}
func minBound(b1, b2 float64) float64 {
	if b1 < b2 {
		return b1
	}
	return b2
}

func NewArea(b1, b2 mgl64.Vec2) Area {
	return Area{
		minX: minBound(b1.X(), b2.X()),
		maxX: maxBound(b1.X(), b2.X()),

		minY: minBound(b1.Y(), b2.Y()),
		maxY: maxBound(b1.Y(), b2.Y()),
	}
}

func (area Area) Vec2Within(vec mgl64.Vec2) bool {
	return vec.X() > area.minX && vec.X() < area.maxX && vec.Y() > area.minY && vec.Y() < area.maxY
}

func (area Area) Vec2WithinOrEqual(vec mgl64.Vec2) bool {
	return vec.X() >= area.minX && vec.X() <= area.maxX && vec.Y() >= area.minY && vec.Y() <= area.maxY
}

func (area Area) Vec3WithinOrEqualXZ(vec mgl64.Vec3) bool {
	return vec.X() >= area.minX && vec.X() <= area.maxX && vec.Z() >= area.minY && vec.Z() <= area.maxY
}

func (area Area) Vec3WithinXZ(vec mgl64.Vec3) bool {
	return vec.X() > area.minX && vec.X() < area.maxX && vec.Z() > area.minY && vec.Z() < area.maxY
}

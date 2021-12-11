package area

import "github.com/go-gl/mathgl/mgl64"

type Area struct {
	min, max mgl64.Vec2
}

func NewArea(min, max mgl64.Vec2) Area {
	return Area{
		min: min,
		max: max,
	}
}

func (area Area) Vec2Within(vec mgl64.Vec2) bool {
	if vec.X() < area.min.X() || vec.X() > area.max.X() {
		return false
	}
	return vec.Y() >= area.min.Y() && vec.Y() <= area.max.Y()
}

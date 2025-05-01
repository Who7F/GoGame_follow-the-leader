package maps

import (
	"follow-the-leader/cmd/camera"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type ColliderProvider interface {
	Draw(screen *ebiten.Image, cam *camera.Camera)
	Collides(x, y, width, height float64) bool
}

type RectColliders struct {
	X, Y     float64
	Layer    int
	Width    float64
	Height   float64
	Rotation float64
	Type     string
	Meta     map[string]string
}

func (c *RectColliders) Draw(screen *ebiten.Image, cam *camera.Camera) {
	vector.StrokeRect(
		screen,
		float32(c.X+cam.X),
		float32(c.Y+cam.Y),
		float32(c.Width),
		float32(c.Height),
		1.0,
		color.Black,
		true,
	)
}

func (c *RectColliders) Collides(x, y, width, height float64) bool {
	return aabb(x, y, width, height, c.X, c.Y, c.Width, c.Height)
}

type CircleColliders struct {
	X, Y   float64
	Layer  int
	Type   string
	Radius float64
	Meta   map[string]string
}

func (c *CircleColliders) Draw(screen *ebiten.Image, cam *camera.Camera) {
	vector.StrokeCircle(
		screen,
		float32(c.X+cam.X),
		float32(c.Y+cam.Y),
		float32(c.Radius),
		1.0,
		color.Black,
		true,
	)
}

func (c *CircleColliders) Collides(x, y, width, height float64) bool {
	return aabb(x, y, width, height, c.X, c.Y, width, height)
}

type PolygonColliders struct {
	X, Y     float64
	Layer    int
	Width    float64
	Height   float64
	Rotation float64
	Type     string
	Meta     map[string]string
	Polygon  []Point
}

func (c *PolygonColliders) Draw(screen *ebiten.Image, cam *camera.Camera) {
	for i := 0; i < len(c.Polygon); i++ {
		p1 := c.Polygon[i]
		p2 := c.Polygon[(i+1)%len(c.Polygon)]
		vector.StrokeLine(
			screen,
			float32(p1.X+c.X+cam.X),
			float32(p1.Y+c.Y+cam.Y),
			float32(p2.X+c.X+cam.X),
			float32(p2.Y+c.Y+cam.Y),
			1.0,
			color.Black,
			true,
		)
	}
}

func (c *PolygonColliders) Collides(x, y, width, height float64) bool {
	return false
}

func aabb(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	return x1 < x2+w2 && x1+w1 > x2 && y1 < y2+h2 && y1+h1 > y2
}

func CircleCollision(cx, cy, x, y, width, height, radius float64) bool {
	// Center of the box
	boxCenterX := x + width/2
	boxCenterY := y + height/2

	// Distance between centers
	dx := boxCenterX - cx
	dy := boxCenterY - cy

	// Clamp the distance to the box
	closestX := clamp(dx, -width/2, width/2)
	closestY := clamp(dy, -height/2, height/2)

	// Find the closest point on the box to the circle
	closestPointX := boxCenterX - closestX
	closestPointY := boxCenterY - closestY

	// Distance from circle center to closest point
	distanceX := closestPointX - cx
	distanceY := closestPointY - cy

	distanceSquared := (distanceX * distanceX) + (distanceY * distanceY)

	return distanceSquared < (radius * radius)
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

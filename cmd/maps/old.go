package maps

type Colliders struct {
	X, Y     float64
	Width    float64
	Height   float64
	Rotation float64
	Type     string
	TileID   int
	Ellipe   bool
	Meta     map[string]string
	Polygon  []Point
}

type ObjectGroup struct {
	Objects []CollisionObject
}

type Point struct {
	X float64
	Y float64
}

type CollisionObject struct {
	ID       int
	Name     string
	Type     string
	X        float64
	Y        float64
	Width    float64
	Height   float64
	Rotation float64

	Ellipe  bool
	Polygon []Point
	//Meta    map[string]string `json:"properties,omitempty"`
}

type Tile struct {
	ID          int
	Image       string
	ImageWidth  int
	ImageHeight int
	ObjectGroup *ObjectGroup
}

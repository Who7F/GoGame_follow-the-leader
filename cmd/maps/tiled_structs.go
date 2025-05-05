package maps

type LayerTiled struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	Type      string        `json:"type"`
	Data      []int         `json:"data,omitempty"`
	Objects   []ObjectTiled `json:"objects,omitempty"`
	Width     int           `json:"width"`
	Height    int           `json:"height"`
	DeawOrder string        `json:"draworder"`
	Class     string        `json:"class,omitempty"`
	Opacity   float64       `json:"opacity"`
	Visible   bool          `json:"visible"`
	X         int           `json:"x"`
	Y         int           `json:"y"`
}

type TilesetTiled struct {
	FirstGID int                 `json:"firstgid"`
	Source   string              `json:"source"`
	Parsed   *TilesetSourceTiled `json:"-"`
}

type TilemapTiled struct {
	Tilesets   []TilesetTiled `json:"tilesets"`
	Tiles      []LayerTiled   `json:"layers"`
	TileWidth  int            `json:"tilewidth"`
	TileHeight int            `json:"tileheight"`
	Width      int            `json:"width"`
	Height     int            `json:"height"`
}

type TilesetSourceTiled struct {
	Class       string      `json:"class,omitempty"`
	Columns     int         `json:"columns"`
	Image       string      `json:"image,omitempty"`
	ImageHeight int         `json:"imageheight"`
	ImageWidth  int         `json:"imagewidth"`
	Margin      int         `json:"margin"`
	Name        string      `json:"name"`
	Spacing     int         `json:"spacing"`
	TileConnt   int         `json:"tilecount"`
	TileHeight  int         `json:"tileheight"`
	TileWidth   int         `json:"tilewidth"`
	Type        string      `json:"tileset"`
	Tiles       []TileTiled `json:"tiles,eomitempty"`
}

type TileTiled struct {
	ID          int    `json:"id"`
	Image       string `json:"image,eomitempty"`
	ImageWidth  int    `json:"imagewidth"`
	ImageHeight int    `json:"imageheight"`
	ObjectGroup *ObjectGroupTiled
}

type ObjectGroupTiled struct {
	Objects []ObjectTiled `json:"objects"`
}

type ObjectTiled struct {
	ID       int          `json:"id"`
	Name     string       `json:"name,omitempty"`
	Type     string       `json:"type"`
	X        float64      `json:"x"`
	Y        float64      `json:"y"`
	Width    float64      `json:"width"`
	Height   float64      `json:"height"`
	Rotation float64      `json:"rotation"`
	Visible  bool         `json:"visible"`
	Ellipe   bool         `json:"ellipse,omitempty"`
	Polygon  []PointTiled `json:"polygon,omitempty"`
	//Meta    map[string]string `json:"properties,omitempty"` Not in use
}

type PointTiled struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

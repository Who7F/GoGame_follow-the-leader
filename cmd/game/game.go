package game

import (
	"fmt"
	"follow-the-leader/cmd/camera"
	"follow-the-leader/cmd/entities"
	"follow-the-leader/cmd/maps"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Game struct holds all entities
type Game struct {
	Player    *entities.Player
	NPCs      []*entities.Npc
	Rums      []*entities.Rum
	Tilemap   *maps.TilemapJSON
	Tilesets  []maps.TileProvider
	Cam       *camera.Camera
	Colliders []*maps.Colliders
}

// New initializes the game
func New() (*Game, error) {
	player, err := entities.NewPlayer(150, 150)
	if err != nil {
		return nil, err
	}

	npcs, err := entities.NewNPCs("assets/entityData/npcs.json")
	if err != nil {
		log.Fatal(err)
	}

	rums, err := entities.NewRums("assets/entityData/rums.json")
	if err != nil {
		log.Fatal(err)
	}

	tilemap, err := maps.NewTilemapJSON("assets/maps/tilesset/startbace.json")
	if err != nil {
		log.Fatal(err)
	}

	tilesets, err := maps.LoadTilesets(tilemap)
	if err != nil {
		log.Fatal(err)
	}

	colliders, err := tilemap.SetColliders(tilesets)
	if err != nil {
		log.Fatal(err)
	}

	camera := camera.NewCamera(0, 0)

	return &Game{
		Player:    player,
		NPCs:      npcs,
		Rums:      rums,
		Tilemap:   tilemap,
		Tilesets:  tilesets,
		Cam:       camera,
		Colliders: colliders,
	}, nil
}

// Update handles game logic
func (g *Game) Update() error {
	g.Player.Update(g.Colliders)

	for _, npc := range g.NPCs {
		npc.Update(g.Player.X, g.Player.Y, g.Colliders)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		targetX, targetY := ebiten.CursorPosition()
		// To call movemnet func
		fmt.Printf("x: %d y: %d", targetX, targetY)
	}
	screenWidth, screenHeight := 640.0, 480.0
	offset := 8.0

	tilemapWidth := float64(g.Tilemap.Tiles[0].Width) * 16.0
	tilemapHeight := float64(g.Tilemap.Tiles[0].Height) * 16.0

	g.Cam.FollowTarget(g.Player.X+offset, g.Player.Y+offset, screenWidth, screenHeight)
	g.Cam.Constrain(tilemapWidth, tilemapHeight, screenWidth, screenHeight)
	dt := 1.0 / 60.0
	g.Player.Anim.Update(dt)
	return nil
}

// Draw renders everything
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0, 0xff, 0xff})

	g.Tilemap.Draw(screen, g.Tilesets, g.Cam)

	ebitenutil.DebugPrint(screen, "Constitution Build!")

	for _, npc := range g.NPCs {
		//	npc.Draw(screen, g.Cam)
		npc.Anim.Draw(screen, g.Cam, npc.X, npc.Y)
	}

	for _, rum := range g.Rums {
		rum.Draw(screen, g.Cam)
	}

	for _, collider := range g.Colliders {
		vector.StrokeRect(
			screen,
			float32(int(collider.X)+int(g.Cam.X)),
			float32(int(collider.Y)+int(g.Cam.Y)),
			float32(collider.Width),
			float32(collider.Height),
			1.0,
			color.Black,
			true,
		)
	}

	g.Player.Anim.Draw(screen, g.Cam, g.Player.X, g.Player.Y)

	//g.Player.Draw(screen, g.Cam)
}

// Layout sets the screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

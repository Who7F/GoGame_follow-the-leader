package game

import (
	"fmt"
	"follow-the-leader/cmd/camera"
	"follow-the-leader/cmd/entities"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Game struct holds all entities
type Game struct {
	Player  *entities.Player
	NPCs    []*entities.Npc
	Rums    []*entities.Rum
	Tilemap *entities.TilemapJSON
	Tileset *entities.Tileset
	Cam     *camera.Camera
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

	tilemap, err := entities.NewTilemapJSON("assets/maps/tilesset/test.json")
	if err != nil {
		log.Fatal(err)
	}

	tileset, err := entities.LoadTileset("assets/maps/tilesset/test.png", 16, 16)
	if err != nil {
		log.Fatal(err)
	}

	camera := camera.NewCamera(0, 0)

	return &Game{
		Player:  player,
		NPCs:    npcs,
		Rums:    rums,
		Tilemap: tilemap,
		Tileset: tileset,
		Cam:     camera,
	}, nil
}

// Update handles game logic
func (g *Game) Update() error {
	g.Player.Update()

	for _, npc := range g.NPCs {
		npc.Update(g.Player.X, g.Player.Y)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		targetX, targetY := ebiten.CursorPosition()
		// To call movemnet func
		fmt.Printf("x: %d y: %d", targetX, targetY)
	}
	camX, camY := 320.0, 240.0

	g.Cam.FollowTarget(g.Player.X, g.Player.Y, camX, camY)

	return nil
}

// Draw renders everything
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0, 0xff, 0xff})

	g.Tilemap.Draw(screen, g.Tileset, g.Cam)

	ebitenutil.DebugPrint(screen, "Constitution Build!")

	for _, npc := range g.NPCs {
		npc.Draw(screen, g.Cam)
	}
	for _, rum := range g.Rums {
		rum.Draw(screen, g.Cam)
	}

	g.Player.Draw(screen, g.Cam)
}

// Layout sets the screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

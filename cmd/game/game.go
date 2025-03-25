package game

import (
	"image/color"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"follow-the-leader/cmd/entities"
)

// Game struct holds all entities
type Game struct {
	Player *entities.Player
	NPCs   []*entities.Npc
	Rums   []*entities.Rum
	//TilemapJSON *entities.TilemapJSON
}

// New initializes the game
func New() (*Game, error) {
	player, err := entities.NewPlayer(150, 150)
	if err != nil {
		return nil, err
	}

	npcs, err := entities.NewNPCs("assets/npcs/npcs.json")
	if err != nil {
		log.Fatal(err) 
	}
	
	rums := entities.NewRums()
	
	//tilemapJSON, err := entities.TilemapJSON("assets/maps/tilesset/test.json")
	//if err != nil {
	//	log.Fatal(err) 
	//}

	return &Game{
		Player: player,
		NPCs:   npcs,
		Rums:   rums,
		//TilemapJSON: tilemapJSON,
	}, nil
}

// Update handles game logic
func (g *Game) Update() error {
	g.Player.Update()

	for _, npc := range g.NPCs {
		npc.Update(g.Player.X, g.Player.Y)
	}

	return nil
}

// Draw renders everything
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0, 0xff, 0xff})
	
	//g.TilemapJSON.Draw(screen)
	
	ebitenutil.DebugPrint(screen, "Constitution Build!")

	g.Player.Draw(screen)
	
	for _, npc := range g.NPCs {
		npc.Draw(screen)
	}
	for _, rum := range g.Rums {
		rum.Draw(screen)
	}
}

// Layout sets the screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

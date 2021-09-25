package game

import (
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/ai"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/unit"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/world"
)

type Game struct {
	World     world.World
	Units     []*unit.Unit
	AIHandler *ai.AIHandler
}

func NewGame(size int) Game {
	return Game{
		World: world.World{
			Tiles: make([][]*world.Tile, size, size),
		},
		AIHandler: initAI(),
	}
}

//here is where we first move, then deal the damage for this tick, starting by structures dealing damage, and then units.
//for fairness, we will randomly choose a structure/unit to deal damage to their target
func (game *Game) tick() {

}

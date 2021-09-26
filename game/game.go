package game

import (
	"sync"

	"github.com/SignedAdam/multiplayer_tower_defense/pkg/ai"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/unit"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/world"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/unit_types"
)

// TODO: optimize concurrency: TBD
type Game struct {
	*sync.RWMutex
	World     *world.World
	Units     []unit.Unit
	AIHandler *ai.AIHandler
}

func NewGame(size int) Game {
	return Game{
		World: &world.World{
			Tiles: make([][]*world.Tile, size, size),
		},
		AIHandler: initAI(),
	}
}

//here is where we first move, then deal the damage for this tick, starting by structures dealing damage, and then units.
//for fairness, we will randomly choose a structure/unit to deal damage to their target
func (game *Game) tick() {

	for _, unit := range game.Units {
		if !unit.IsDead() {
			unit.Act(game.Units, game.World)
		}

	}
}

func (game *Game) CreateUnit(unitTypeID, x, y int) (unit.Unit, error) {
	game.Lock()
	defer game.Unlock()

	unitType, err := unit_types.GetUnitType(unitTypeID)
	if err != nil {
		return nil, err
	}

	newUnit, err := unitType.CreateUnit(x, y)
	if err != nil {
		return nil, err
	}

	return newUnit, nil
}

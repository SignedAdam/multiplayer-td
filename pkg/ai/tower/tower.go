package tower

import (
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/ai"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/unit"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/world"
)

type towerAI struct {
}

func NewTowerAI() ai.AI {
	return &towerAI{}
}

func (ai *towerAI) Move(world *world.World, target unit.Unit) error {
	return nil
}

func (ai *towerAI) ChooseTarget(world *world.World, units []unit.Unit) unit.Unit {

	return nil
}

func (ai *towerAI) Placeholder() {

}

package npc

import (
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/ai"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/unit"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/world"
)

type npcAI struct {
}

func NewNPCAI() ai.AI {
	return &npcAI{}
}

func (ai *npcAI) Move(world *world.World, target unit.Unit) error {
	return nil
}

func (ai *npcAI) ChooseTarget(world *world.World, units []unit.Unit) unit.Unit {

	return nil
}

func (ai *npcAI) Placeholder() {

}

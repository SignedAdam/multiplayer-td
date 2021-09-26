package npc

import (
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/ai"
	unit "github.com/SignedAdam/multiplayer_tower_defense/pkg/game/unit"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/world"
	"github.com/google/uuid"
)

type npc struct {
	id        uuid.UUID
	npcTypeID int
	x         int
	y         int
	health    int
	damage    int
	target    unit.Unit
	ai        ai.AI
}

func NewNPC(npcTypeID, x, y, health, damage int, AI ai.AI) unit.Unit {
	return &npc{
		id:        uuid.New(),
		npcTypeID: npcTypeID,
		x:         x,
		y:         y,
		health:    health,
		damage:    damage,
		ai:        AI,
	}
}

func (npc *npc) Act(units []unit.Unit, world *world.World) {
	if npc.target.TargetDead() {
		npc.target = chooseNewTarget(units)
	}

	err := npc.ai.Move(world, npc.target)
	if err != nil {
		// TODO: handle
	}

	npc.target.Hit(npc.damage)

}

// TODO
func chooseNewTarget(units []unit.Unit) unit.Unit {
	return nil
}

func (npc *npc) TargetDead() bool {
	return npc.target.IsDead() == true
}

func (npc *npc) IsDead() bool {
	return npc.health <= 0
}

func (npc *npc) GetPos() (int, int) {
	return npc.x, npc.y
}

func (npc *npc) GetHealth() int {
	return npc.health
}

func (npc *npc) GetType() int {
	return npc.npcTypeID
}

func (npc *npc) Hit(damage int) {
	npc.health -= damage
}

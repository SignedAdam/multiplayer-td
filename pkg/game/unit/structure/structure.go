package structure

import (
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/ai"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/unit"
	"github.com/google/uuid"
)

type structure struct {
	id      uuid.UUID
	npcType int
	x       int
	y       int
	health  int
	damage  int
	target  unit.Unit
	ai      ai.AI
}

func NewStructure(npcType, x, y, health, damage int, AI ai.AI) unit.Unit {
	return &structure{
		id:      uuid.New(),
		npcType: npcType,
		x:       x,
		y:       y,
		health:  health,
		damage:  damage,
		ai:      AI,
	}
}

func (structure *structure) Act(units []*unit.Unit) {
	if structure.target.TargetDead() {
		structure.target = chooseNewTarget(units)
	}

	//attack
	structure.target.Hit(structure.damage)

}

// TODO
func chooseNewTarget(units []*unit.Unit) unit.Unit {
	return nil
}

func (structure *structure) TargetDead() bool {
	return structure.target.IsDead() == true
}

func (structure *structure) IsDead() bool {
	return structure.health <= 0
}

func (structure *structure) GetPos() (int, int) {
	return structure.x, structure.y
}

func (structure *structure) GetHealth() int {
	return structure.health
}

func (structure *structure) GetType() int {
	return structure.npcType
}

func (structure *structure) Hit(damage int) {
	structure.health -= damage
}

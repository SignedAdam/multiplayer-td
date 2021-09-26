package tower

import (
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/ai"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/unit"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/world"
	"github.com/google/uuid"
)

type tower struct {
	id      uuid.UUID
	npcType int
	x       int
	y       int
	health  int
	damage  int
	target  unit.Unit
	ai      ai.AI
}

func NewTower(npcType, x, y, health, damage int, AI ai.AI) unit.Unit {
	return &tower{
		id:      uuid.New(),
		npcType: npcType,
		x:       x,
		y:       y,
		health:  health,
		damage:  damage,
		ai:      AI,
	}
}

func (tower *tower) Act(units []unit.Unit, world *world.World) {
	if tower.target.TargetDead() {
		tower.target = tower.ai.ChooseTarget(world, units)
	}

	tower.target.Hit(tower.damage)

}

func (tower *tower) TargetDead() bool {
	return tower.target.IsDead() == true
}

func (tower *tower) IsDead() bool {
	return tower.health <= 0
}

func (tower *tower) GetPos() (int, int) {
	return tower.x, tower.y
}

func (tower *tower) GetHealth() int {
	return tower.health
}

func (tower *tower) GetType() int {
	return tower.npcType
}

func (tower *tower) Hit(damage int) {
	tower.health -= damage
}

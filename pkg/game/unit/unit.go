package unit

import "github.com/SignedAdam/multiplayer_tower_defense/pkg/game/world"

type Unit interface {
	Act(units []Unit, world *world.World)
	TargetDead() bool
	IsDead() bool
	GetPos() (int, int)
	GetHealth() int
	GetType() int
	Hit(damage int)
}

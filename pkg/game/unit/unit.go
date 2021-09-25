package unit

type Unit interface {
	Act(units []*Unit)
	TargetDead() bool
	IsDead() bool
	GetPos() (int, int)
	GetHealth() int
	GetType() int
	Hit(damage int)
}

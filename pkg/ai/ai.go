package ai

import (
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/unit"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/game/world"
)

type AI interface {
	Placeholder()
	Move(world *world.World, target unit.Unit) error
	ChooseTarget(world *world.World, units []unit.Unit) unit.Unit
	//ChooseNextTarget
}

type AIHandler struct {
	ai map[int]AI
}

func NewAIHandler() *AIHandler {
	return &AIHandler{
		ai: make(map[int]AI),
	}
}

func (handler *AIHandler) AddAI(unitType int, AILogic AI) {
	handler.ai[unitType] = AILogic
}

func (handler *AIHandler) GetAI(unitType int) AI {
	return handler.ai[unitType]
}

package game

import (
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/ai"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/ai/npc"
	"github.com/SignedAdam/multiplayer_tower_defense/pkg/ai/tower"
)

func initGame() {

}

func initAI() *ai.AIHandler {
	AIHandler := ai.NewAIHandler()
	AIHandler.AddAI(0, tower.NewTowerAI())
	AIHandler.AddAI(1, npc.NewNPCAI())

	return AIHandler
}

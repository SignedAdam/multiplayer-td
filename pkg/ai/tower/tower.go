package tower

import "github.com/SignedAdam/multiplayer_tower_defense/pkg/ai"

type towerAI struct {
}

func NewTowerAI() ai.AI {
	return &towerAI{}
}

func (ai *towerAI) Placeholder() {

}

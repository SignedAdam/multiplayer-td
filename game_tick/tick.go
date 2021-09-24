package game_tick

import (
	"github.com/SignedAdam/multiplayer_tower_defense/game"
)

type ticker struct {
	game     *game.Game
	tickRate int
}

func newTicker(game *game.Game, tickRate int) {
	ticker := &ticker{
		game:     game,
		tickRate: tickRate,
	}
	go ticker.startTicker()
}

func (t *ticker) startTicker() {

}

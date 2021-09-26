package game_tick

import (
	"log"
	"math"
	"os"
	"time"

	"github.com/SignedAdam/multiplayer_tower_defense/game"
)

type ticker struct {
	sigStopServerChan chan os.Signal
	game              *game.Game
	tickPeriod        int64
	logger            *log.Logger
}

func NewTicker(game *game.Game, tickRate int, logger *log.Logger, sigStopServerChan chan os.Signal) {
	ticker := &ticker{
		game:              game,
		tickPeriod:        int64(math.Round(1000 / float64(tickRate))),
		sigStopServerChan: sigStopServerChan,
		logger:            logger,
	}
	go ticker.startTicker()
}

func (t *ticker) startTicker() {
	for {
		select {
		// when the sigStopServerChan channel is signalled then we return from the function to kill the thread
		case <-t.sigStopServerChan:
			t.logger.Println("Exiting thread")
			return
		default:
			tickStartTime := time.Now()

			t.game.Lock()
			t.tick()
			t.game.Unlock()

			millisPassed := time.Now().Sub(tickStartTime).Milliseconds()
			if millisPassed < int64(t.tickPeriod) {
				time.Sleep((time.Duration(t.tickPeriod - millisPassed)) * time.Millisecond)
			}
		}
	}
}
func (t *ticker) tick() {
	t.logger.Printf("tick %v", time.Now().UnixMilli())
}

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SignedAdam/multiplayer_tower_defense/game"
	"github.com/SignedAdam/multiplayer_tower_defense/game_tick"
)

func main() {
	sigStopServerChan := make(chan os.Signal, 1)
	signal.Notify(sigStopServerChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	game := game.NewGame(10)

	game_tick.NewTicker(&game, 1, log.Default(), sigStopServerChan)

	<-sigStopServerChan
}

func initAI() {

}

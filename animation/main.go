package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"../../../Desktop/game_go/GameProgramming_Go/animation"
	"../../../Desktop/game_go/GameProgramming_Go/animation"
	"../../../Desktop/game_go/GameProgramming_Go/animation"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	scenemanager.SetScene(&scenes.StartScene{})

	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 2.0, "Run game")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
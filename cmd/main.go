package main

import (
	"github.com/gen2brain/raylib-go/raylib"

	"gitea.theedgeofrage.com/theedgeofrage/yeetris/game"
)

func main() {
	game := game.Game{}
	game.Init()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
	}

	rl.CloseWindow()
}

package game

import (
	"gitea.theedgeofrage.com/theedgeofrage/yeetris/elements"
	"github.com/gen2brain/raylib-go/raylib"
)

const DEBUG = false

// Game type
type Game struct {
	FramesCounter int32
	TickRate      int32
	GameOver      bool
	Pause         bool

	Scale rl.Vector2

	Board *Board
	UI    *UI
}

// Init - Initialize game
func (g *Game) Init() {
	g.Scale = rl.NewVector2(40, 40)

	g.FramesCounter = 0
	g.TickRate = 30
	g.GameOver = false
	g.Pause = false
	g.Board = InitBoard()
	g.UI = InitUI(g.Scale)
}

func (g *Game) Update() {
	defer g.Board.ClearLines()
	if g.GameOver {
		if rl.IsKeyPressed(rl.KeyEnter) {
			g.Init()
			g.GameOver = false
		}
		return
	}
	if rl.IsKeyPressed(rl.KeyP) {
		g.Pause = !g.Pause
	}

	if g.Pause {
		return
	}

	if rl.IsKeyPressed(rl.KeyC) {
		g.Board.HoldPiece()
	}
	if rl.IsKeyPressed(rl.KeyI) {
		g.Board.MoveActivePiece(1)
	}
	if rl.IsKeyPressed(rl.KeyH) {
		g.Board.MoveActivePiece(-1)
	}
	if rl.IsKeyPressed(rl.KeyE) {
		g.Board.RotateActivePiece(true)
	}
	if rl.IsKeyPressed(rl.KeyN) {
		g.Board.DescendActivePiece()
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		g.Board.DropActivePiece()
	}
	if g.FramesCounter%g.TickRate == 0 && !DEBUG {
		g.Board.DescendActivePiece()
	}
	g.FramesCounter++
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(elements.Black)

	g.UI.Draw(g)

	if !g.GameOver {
		g.Board.DrawBoard(g.UI.BoardOffset, g.Scale)
		g.Board.DrawNextPiece(g.UI.NextPiecePosition, g.Scale)
		g.Board.DrawHeldPiece(g.UI.HeldPiecePosition, g.Scale)
	} else {
		rl.DrawText(
			"PRESS [ENTER] TO PLAY AGAIN",
			int32(rl.GetScreenWidth())/2-rl.MeasureText("PRESS [ENTER] TO PLAY AGAIN", 20)/2,
			int32(rl.GetScreenHeight())/2-50,
			20,
			elements.LightGray,
		)
	}

	rl.EndDrawing()
}

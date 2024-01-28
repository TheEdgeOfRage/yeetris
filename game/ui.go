package game

import (
	"fmt"

	"gitea.theedgeofrage.com/theedgeofrage/yeetris/elements"
	"github.com/gen2brain/raylib-go/raylib"
)

type UI struct {
	ScreenWidth       int32
	ScreenHeight      int32
	BoardWidth        int32
	BoardHeight       int32
	BoardOffset       rl.Vector2
	UIOffsetWidth     int32
	NextPiecePosition rl.Vector2
	HeldPiecePosition rl.Vector2

	Score int
}

func InitUI(scale rl.Vector2) *UI {
	boardWidth := 10 * int32(scale.X)
	boardHeight := 20 * int32(scale.Y)
	screenWidth := boardWidth + 40 + 300
	screenHeight := boardHeight + 40
	rl.InitWindow(screenWidth, screenHeight, "Yeetris")

	return &UI{
		ScreenWidth:       screenWidth,
		ScreenHeight:      screenHeight,
		BoardWidth:        boardWidth,
		BoardHeight:       boardHeight,
		BoardOffset:       rl.NewVector2(20, 20),
		UIOffsetWidth:     boardWidth + 40 + 20,
		NextPiecePosition: rl.NewVector2(float32(boardWidth+40+20+110), 250),
		HeldPiecePosition: rl.NewVector2(float32(boardWidth+40+20+110), 510),
	}
}

func drawRectangleBorder(x, y, width, height, border int32, color rl.Color) {
	rl.DrawRectangle(x-border, y-border, width+2*border, height+2*border, color)
	rl.DrawRectangle(x, y, width, height, elements.Black)
}

func (u *UI) Draw(g *Game) {
	drawRectangleBorder(int32(u.BoardOffset.X), int32(u.BoardOffset.Y), u.BoardWidth, u.BoardHeight, 20, elements.Gray)
	rl.DrawText("Yeetris", u.UIOffsetWidth, 20, 40, rl.White)
	rl.DrawText(fmt.Sprintf("Score: %d", u.Score), u.UIOffsetWidth, 70, 20, rl.White)
	rl.DrawText("Next Piece:", u.UIOffsetWidth, 110, 20, rl.White)
	drawRectangleBorder(u.UIOffsetWidth+10, 150, 200, 200, 10, elements.Gray)
	rl.DrawText("Hold:", u.UIOffsetWidth, 370, 20, rl.White)
	drawRectangleBorder(u.UIOffsetWidth+10, 410, 200, 200, 10, elements.Gray)

	if g.Pause {
		rl.DrawText(
			"GAME PAUSED",
			u.ScreenWidth/2-rl.MeasureText("GAME PAUSED", 40)/2,
			u.ScreenHeight/2-40,
			40,
			elements.LightGray,
		)
	}
}

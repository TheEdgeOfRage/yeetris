package elements

import "github.com/gen2brain/raylib-go/raylib"

type Block struct {
	Position rl.Vector2
	Color    rl.Color
	LongBoi  bool
}

func (b *Block) Draw(piecePosition rl.Vector2, scale rl.Vector2, offset rl.Vector2) {
	basePos := rl.Vector2Add(b.Position, piecePosition)
	if basePos.Y < 0 {
		return
	}
	pos := rl.Vector2Add(rl.Vector2Multiply(basePos, scale), offset)
	rl.DrawRectangleV(pos, scale, b.Color)
}

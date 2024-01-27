package elements

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Piece struct {
	Blocks []*Block
}

var Pieces = []*Piece{
	{
		Blocks: []*Block{
			{rl.NewVector2(-1, 0), Cyan, false},
			{rl.NewVector2(0, 0), Cyan, false},
			{rl.NewVector2(1, 0), Cyan, true},
			{rl.NewVector2(2, 0), Cyan, false},
		},
	},
	{
		Blocks: []*Block{
			{rl.NewVector2(-1, 0), Blue, false},
			{rl.NewVector2(-1, -1), Blue, false},
			{rl.NewVector2(0, 0), Blue, false},
			{rl.NewVector2(1, 0), Blue, false},
		},
	},
	{
		Blocks: []*Block{
			{rl.NewVector2(-1, 0), Orange, false},
			{rl.NewVector2(0, 0), Orange, false},
			{rl.NewVector2(1, 0), Orange, false},
			{rl.NewVector2(1, -1), Orange, false},
		},
	},
	{
		Blocks: []*Block{
			{rl.NewVector2(-1, 0), Purple, false},
			{rl.NewVector2(0, 0), Purple, false},
			{rl.NewVector2(1, 0), Purple, false},
			{rl.NewVector2(0, -1), Purple, false},
		},
	},
	{
		Blocks: []*Block{
			{rl.NewVector2(-1, -1), Red, false},
			{rl.NewVector2(0, -1), Red, false},
			{rl.NewVector2(0, 0), Red, false},
			{rl.NewVector2(1, 0), Red, false},
		},
	},
	{
		Blocks: []*Block{
			{rl.NewVector2(-1, 0), Green, false},
			{rl.NewVector2(0, 0), Green, false},
			{rl.NewVector2(0, -1), Green, false},
			{rl.NewVector2(1, -1), Green, false},
		},
	},
	{
		Blocks: []*Block{
			{rl.NewVector2(0, -1), Yellow, false},
			{rl.NewVector2(1, -1), Yellow, false},
			{rl.NewVector2(0, 0), Yellow, false},
			{rl.NewVector2(1, 0), Yellow, false},
		},
	},
}

func (p *Piece) Draw(pos rl.Vector2, scale rl.Vector2, offset rl.Vector2) {
	for _, block := range p.Blocks {
		block.Draw(pos, scale, offset)
	}
}

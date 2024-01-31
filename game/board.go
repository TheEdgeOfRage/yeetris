package game

import (
	"math"

	"gitea.theedgeofrage.com/theedgeofrage/yeetris/elements"
	"github.com/gen2brain/raylib-go/raylib"
)

type Board struct {
	Board [10][20]*elements.Block

	ActivePiece    *elements.Piece
	ActivePiecePos rl.Vector2
	NextPiece      *elements.Piece
	HeldPiece      *elements.Piece
	AlreadyHeld    bool
}

func InitBoard() *Board {
	board := &Board{
		NextPiece: elements.GetRandomPiece(),
	}
	_ = board.getNextPiece()

	return board
}

func (b *Board) getNextPiece() error {
	b.AlreadyHeld = false
	b.ActivePiecePos = rl.NewVector2(5, 0)
	b.ActivePiece = b.NextPiece
	b.NextPiece = elements.GetRandomPiece()
	for _, block := range b.ActivePiece.Blocks {
		newX, newY := b.ActivePiecePos.X+block.Position.X, b.ActivePiecePos.Y+block.Position.Y
		if newY >= 0 && b.Board[int(newX)][int(newY)] != nil {
			return ErrGameOver
		}
	}
	return nil
}

func (b *Board) setBlocksFromActivePiece() error {
	for _, block := range b.ActivePiece.Blocks {
		newX, newY := b.ActivePiecePos.X+block.Position.X, b.ActivePiecePos.Y+block.Position.Y
		if newY < 0 {
			return ErrGameOver
		}
		b.Board[int(newX)][int(newY)] = &elements.Block{
			Position: rl.NewVector2(newX, newY),
			Color:    b.ActivePiece.Blocks[0].Color,
		}
	}

	return nil
}

func (b *Board) checkHorizontalCollision(left bool) bool {
	for _, activeBlock := range b.ActivePiece.Blocks {
		xPos := int(b.ActivePiecePos.X + activeBlock.Position.X)
		yPos := int(b.ActivePiecePos.Y + activeBlock.Position.Y)
		if left {
			if xPos == 0 {
				return true
			}
			if yPos >= 0 && b.Board[xPos-1][yPos] != nil {
				return true
			}
		} else {
			if xPos == 9 {
				return true
			}
			if yPos >= 0 && b.Board[xPos+1][yPos] != nil {
				return true
			}
		}
	}
	return false
}

func (b *Board) descendActivePiece() (bool, error) {
	b.ActivePiecePos.Y += 1
	for _, activeBlock := range b.ActivePiece.Blocks {
		if b.ActivePiecePos.Y+activeBlock.Position.Y == 20 ||
			b.Board[int(b.ActivePiecePos.X+activeBlock.Position.X)][int(b.ActivePiecePos.Y+activeBlock.Position.Y)] != nil {
			b.ActivePiecePos.Y -= 1
			return true, b.setBlocksFromActivePiece()
		}
	}
	return false, nil
}

// DescendActivePiece moves the piece down by one block. If the piece cannot move down, it will be placed on the board.
// After that, a new piece will be generated. If the new piece cannot be placed, a game over error is returned.
func (b *Board) DescendActivePiece() error {
	placed, err := b.descendActivePiece()
	if err != nil {
		return err
	}

	if placed {
		return b.getNextPiece()
	}
	return nil
}

// DropActivePiece moves the piece down until it cannot move down anymore. After that, a new piece will be generated.
// If the new piece cannot be placed, a game over error is returned.
func (b *Board) DropActivePiece() error {
	for {
		placed, err := b.descendActivePiece()
		if err != nil {
			return err
		}
		if placed {
			break
		}
	}
	return b.getNextPiece()
}

// MoveActivePiece moves the piece left or right. If the piece cannot move in the specified direction, nothing happens.
func (b *Board) MoveActivePiece(direction int) {
	if !b.checkHorizontalCollision(direction == -1) {
		b.ActivePiecePos.X += float32(direction)
	}
}

func getCollisionDepth(collisionDepth int, longBoi bool) int {
	if longBoi || collisionDepth == 2 {
		return 2
	} else {
		return 1
	}
}

// RotateActivePiece rotates the piece clockwise or counter-clockwise. If the piece cannot rotate in the specified
// direction, nothing happens.
func (b *Board) RotateActivePiece(clockwise bool) {
	if b.ActivePiece.Blocks[0].Color == elements.Yellow {
		return
	}

	angle := rl.Deg2rad * 90
	if !clockwise {
		angle = -angle
	}

	tmpPiece := &elements.Piece{
		Blocks: make([]*elements.Block, 4),
	}
	for i := 0; i < 4; i++ {
		newPos := rl.Vector2Rotate(b.ActivePiece.Blocks[i].Position, float32(angle))
		newPos.X = float32(math.Round(float64(newPos.X)))
		newPos.Y = float32(math.Round(float64(newPos.Y)))
		tmpPiece.Blocks[i] = &elements.Block{
			Position: newPos,
			Color:    b.ActivePiece.Blocks[i].Color,
			LongBoi:  b.ActivePiece.Blocks[i].LongBoi,
		}
	}
	collisionLeft := 0
	collisionRight := 0
	collisionBottom := 0
	for _, block := range tmpPiece.Blocks {
		newX, newY := b.ActivePiecePos.X+block.Position.X, b.ActivePiecePos.Y+block.Position.Y
		if newY < 0 {
			continue
		}
		if newX < 0 && collisionLeft == 0 {
			collisionLeft = getCollisionDepth(collisionLeft, block.LongBoi)
		}
		if newX > 9 && collisionRight == 0 {
			collisionRight = getCollisionDepth(collisionRight, block.LongBoi)
		}
		if newY > 19 && collisionBottom == 0 {
			collisionBottom = getCollisionDepth(collisionBottom, block.LongBoi)
		}
		if collisionLeft != 0 || collisionRight != 0 || collisionBottom != 0 {
			continue
		}
		if b.Board[int(newX)][int(newY)] != nil {
			if newX < b.ActivePiecePos.X && collisionLeft == 0 {
				collisionLeft = getCollisionDepth(collisionLeft, block.LongBoi)
			}
			if newX > b.ActivePiecePos.X && collisionRight == 0 {
				collisionRight = getCollisionDepth(collisionRight, block.LongBoi)
			}
			if newY > b.ActivePiecePos.Y && collisionBottom == 0 {
				collisionBottom = getCollisionDepth(collisionBottom, block.LongBoi)
			}
		}
	}
	if collisionLeft != 0 && collisionRight != 0 {
		return
	}

	b.ActivePiecePos.Y -= float32(collisionBottom)
	b.ActivePiecePos.X += float32(collisionLeft)
	b.ActivePiecePos.X -= float32(collisionRight)
	b.ActivePiece = tmpPiece
}

func (b *Board) HoldPiece() {
	if b.AlreadyHeld {
		return
	}
	b.AlreadyHeld = true
	if b.HeldPiece == nil {
		b.HeldPiece = b.ActivePiece
		_ = b.getNextPiece()
		return
	}
	b.HeldPiece, b.ActivePiece = b.ActivePiece, b.HeldPiece
	b.ActivePiecePos = rl.NewVector2(5, 0)
}

// ClearLines clears all lines that are completely filled.
func (g *Board) ClearLines() int {
	clearedLines := 0
	for y := 19; y >= 0; {
		skip := false
		for x := 0; x < 10; x++ {
			if g.Board[x][y] == nil {
				skip = true
				break
			}
		}
		if skip {
			y--
			continue
		}
		clearedLines++
		for ty := y; ty >= 0; ty-- {
			for x := 0; x < 10; x++ {
				// always clear top line
				if ty == 0 {
					g.Board[x][ty] = nil
					continue
				}
				// move all blocks above cleared line down
				if g.Board[x][ty-1] != nil {
					g.Board[x][ty] = g.Board[x][ty-1]
					g.Board[x][ty].Position.Y++
				} else {
					g.Board[x][ty] = nil
				}
			}
		}
	}
	return clearedLines
}

func (b *Board) DrawGrid(offset, scale rl.Vector2) {
	for i := 0; i < 10; i++ {
		rl.DrawLine(
			int32(scale.X*float32(i)+offset.X),
			int32(offset.Y),
			int32(scale.X*float32(i)+offset.X),
			int32(20*scale.Y+offset.Y),
			elements.Gray,
		)
	}
	for i := 0; i < 20; i++ {
		rl.DrawLine(
			int32(offset.X),
			int32(scale.Y*float32(i)+offset.Y),
			int32(10*scale.X+offset.X),
			int32(scale.Y*float32(i)+offset.Y),
			elements.Gray,
		)
	}
}

func (b *Board) drawUIPiece(piece *elements.Piece, position, scale rl.Vector2) {
	centerOffset := rl.NewVector2(0, 0)
	for _, block := range piece.Blocks {
		centerOffset = rl.Vector2Add(
			centerOffset,
			rl.Vector2Add(block.Position, rl.NewVector2(0.5, 0.5)),
		)
	}
	centerOffset = rl.Vector2Divide(centerOffset, rl.NewVector2(4, 4))
	centerOffset = rl.Vector2Add(centerOffset, rl.NewVector2(0, 1))
	centerOffset = rl.Vector2Multiply(centerOffset, scale)
	for _, block := range piece.Blocks {
		block.Draw(rl.NewVector2(0, 1), scale, rl.Vector2Subtract(position, centerOffset))
	}
}

// DrawNextPiece draws the next piece in the UI at the specified position.
func (b *Board) DrawNextPiece(position, scale rl.Vector2) {
	b.drawUIPiece(b.NextPiece, position, scale)
}

// DrawHeldPiece draws the held piece in the UI at the specified position.
func (b *Board) DrawHeldPiece(position, scale rl.Vector2) {
	if b.HeldPiece == nil {
		return
	}
	b.drawUIPiece(b.HeldPiece, position, scale)
}

// DrawBoard draws the board with placed blocks and the active piece.
func (b *Board) DrawBoard(offset rl.Vector2, scale rl.Vector2) {
	// b.DrawGrid(offset, scale)
	for i := 0; i < 20; i++ {
		for j := 0; j < 10; j++ {
			if b.Board[j][i] == nil {
				continue
			}

			b.Board[j][i].Draw(rl.NewVector2(0, 0), scale, offset)
		}
	}
	b.ActivePiece.Draw(b.ActivePiecePos, scale, offset)
}

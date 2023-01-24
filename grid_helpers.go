package golang_ebiten_helpers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

func DrawGrid(screen *ebiten.Image, width, height, cellSize uint, lineColor color.RGBA) {
	rows := height / cellSize
	cols := width / cellSize

	curRow := uint(1)
	curCol := uint(1)

	for curRow <= rows || curCol <= cols {
		if curRow <= rows {
			y := float64(curRow * cellSize)
			ebitenutil.DrawLine(screen, 0, y, float64(cols*cellSize), y, lineColor)
			curRow++
		}
		if curCol <= cols {
			x := float64(curCol * cellSize)
			ebitenutil.DrawLine(screen, x, 0, x, float64(rows*cellSize), lineColor)
			curCol++
		}
	}
}

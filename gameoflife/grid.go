package gameoflife

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	Cells [][]Cell
	Rows  int32
	Cols  int32
}

func (g *Grid) Update() {
	nextState := make([][]Cell, g.Rows)
	for rowIndex := int32(0); rowIndex < g.Rows; rowIndex++ {
		nextState[rowIndex] = make([]Cell, g.Cols)

		for colIndex := int32(0); colIndex < g.Cols; colIndex++ {
			cell := &g.Cells[rowIndex][colIndex]

			nextState[rowIndex][colIndex] = Cell{
				Row:    rowIndex,
				Col:    colIndex,
				Status: cell.Update(g),
			}
		}
	}

	g.Cells = nextState
}

func (g *Grid) Draw(cellSize int32) {
	g.WalkCells(func(grid *Grid, cell *Cell) {
		if cell.IsAlive() {
			cell.Fill(cellSize, rl.White)
		}
		cell.Outline(cellSize, rl.Gray)
	})
}

func (g *Grid) CellAt(row int32, col int32) (*Cell, error) {
	if row < 0 || row >= int32(len(g.Cells)) || col < 0 || col >= int32(len(g.Cells[row])) {
		return &Cell{Status: Dead}, fmt.Errorf("index out of bounds: row %d, col %d", row, col)
	}

	return &g.Cells[row][col], nil
}

func (g *Grid) WalkCells(callback func(grid *Grid, cell *Cell)) {
	for rowIndex, row := range g.Cells {
		for colIndex := range row {
			callback(g, &g.Cells[rowIndex][colIndex])
		}
	}
}

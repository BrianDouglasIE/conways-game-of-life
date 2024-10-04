package main

import (
	"main/gameoflife"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 800
	COLS          = 5
	ROWS          = 5
	CELL_SIZE     = 160
)

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Conway's Game of Life")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	initialCells := make([][]gameoflife.Cell, ROWS)
	for rowIndex := int32(0); rowIndex < ROWS; rowIndex++ {
		initialCells[rowIndex] = make([]gameoflife.Cell, COLS)

		for colIndex := int32(0); colIndex < COLS; colIndex++ {
			status := gameoflife.Dead
			if rowIndex == 1 && colIndex == 2 ||
				rowIndex == 2 && colIndex == 2 ||
				rowIndex == 3 && colIndex == 2 {
				status = gameoflife.Alive
			}
			initialCells[rowIndex][colIndex] = gameoflife.Cell{Row: rowIndex, Col: colIndex, Status: status}
		}
	}

	grid := gameoflife.Grid{Rows: ROWS, Cols: COLS, Cells: initialCells}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if rl.IsKeyPressed(rl.KeyUp) {
			grid.Update()
		}

		grid.Draw(CELL_SIZE)

		rl.EndDrawing()
	}
}

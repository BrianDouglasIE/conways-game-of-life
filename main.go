package main

import (
	"main/gameoflife"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 800
	COLS          = 16
	ROWS          = 16
	CELL_SIZE     = SCREEN_WIDTH / COLS
)

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Conway's Game of Life")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	initialCells := make([][]gameoflife.Cell, ROWS)
	for rowIndex := int32(0); rowIndex < ROWS; rowIndex++ {
		initialCells[rowIndex] = make([]gameoflife.Cell, COLS)

		for colIndex := int32(0); colIndex < COLS; colIndex++ {
			initialCells[rowIndex][colIndex] = gameoflife.Cell{Row: rowIndex, Col: colIndex, Status: gameoflife.Dead}
		}
	}

	grid := gameoflife.Grid{Rows: ROWS, Cols: COLS, Cells: initialCells}

	for !rl.WindowShouldClose() {
		mousePoint := rl.GetMousePosition()
		grid.WalkCells(func(grid *gameoflife.Grid, cell *gameoflife.Cell) {
			cellRect := rl.Rectangle{Width: CELL_SIZE, Height: CELL_SIZE, X: float32(cell.Col) * CELL_SIZE, Y: float32(cell.Row) * CELL_SIZE}
			if rl.CheckCollisionPointRec(mousePoint, cellRect) {
				if cell.Status == gameoflife.Dead {
					cell.Fill(CELL_SIZE, rl.Green)
				}

				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					if cell.Status == gameoflife.Alive {
						cell.Status = gameoflife.Dead
					} else if cell.Status == gameoflife.Dead {
						cell.Status = gameoflife.Alive
					}
				}
			}
		})

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if rl.IsKeyPressed(rl.KeyUp) {
			grid.Update()
		}

		grid.Draw(CELL_SIZE)

		rl.EndDrawing()
	}
}

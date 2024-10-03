package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 800
	COLUMNS       = 5
	ROWS          = 5
	CELL_SIZE     = 160
)

type GameGrid [ROWS][COLUMNS]Cell

type CellStatus int32

const (
	Alive CellStatus = iota
	Dead
)

type Cell struct {
	Row           int32
	Col           int32
	Status        CellStatus
	NeighbourHood NeighbourHood
}

func (c *Cell) ShouldDie() bool {
	return c.Status == Alive && (c.NeighbourHood.IsOverPopulated() || c.NeighbourHood.IsUnderPopulated())
}

func (c *Cell) ShouldLive() bool {
	return c.Status == Dead && c.NeighbourHood.IsOptimal()
}

func (c *Cell) StatusOfCellAbove(grid *GameGrid) (CellStatus, error) {
	if c.Row < 1 {
		return Dead, fmt.Errorf("No cells above row: %v, col: %v", c.Row, c.Col)
	}
	return grid[c.Row-1][c.Col].Status, nil
}

func (c *Cell) StatusOfCellBelow(grid *GameGrid) (CellStatus, error) {
	if c.Row >= ROWS-1 {
		return Dead, fmt.Errorf("No cells below row: %v, col: %v", c.Row, c.Col)
	}
	return grid[c.Row+1][c.Col].Status, nil
}

func (c *Cell) StatusOfCellToLeft(grid *GameGrid) (CellStatus, error) {
	if c.Col < 1 {
		return Dead, fmt.Errorf("No cells left of row: %v, col: %v", c.Row, c.Col)
	}
	return grid[c.Row][c.Col-1].Status, nil
}

func (c *Cell) StatusOfCellToRight(grid *GameGrid) (CellStatus, error) {
	if c.Col >= COLUMNS-1 {
		return Dead, fmt.Errorf("No cells right of row: %v, col: %v", c.Row, c.Col)
	}
	return grid[c.Row][c.Col+1].Status, nil
}

func (c *Cell) StatusOfCellAboveLeft(grid *GameGrid) (CellStatus, error) {
	if c.Row < 1 || c.Col < 1 {
		return Dead, fmt.Errorf("No cells above left of row: %v, col: %v", c.Row, c.Col)
	}
	return grid[c.Row-1][c.Col-1].Status, nil
}

func (c *Cell) StatusOfCellAboveRight(grid *GameGrid) (CellStatus, error) {
	if c.Row < 1 || c.Col >= COLUMNS-1 {
		return Dead, fmt.Errorf("No cells above right of row: %v, col: %v", c.Row, c.Col)
	}
	return grid[c.Row-1][c.Col+1].Status, nil
}

func (c *Cell) StatusOfCellBelowLeft(grid *GameGrid) (CellStatus, error) {
	if c.Row >= ROWS-1 || c.Col < 1 {
		return Dead, fmt.Errorf("No cells below left of row: %v, col: %v", c.Row, c.Col)
	}
	return grid[c.Row+1][c.Col-1].Status, nil
}

func (c *Cell) StatusOfCellBelowRight(grid *GameGrid) (CellStatus, error) {
	if c.Row >= ROWS-1 || c.Col >= COLUMNS-1 {
		return Dead, fmt.Errorf("No cells below right of row: %v, col: %v", c.Row, c.Col)
	}
	return grid[c.Row+1][c.Col+1].Status, nil
}

func (c *Cell) UpdateNeighbourHood(grid *GameGrid) {
	c.NeighbourHood.Above, _ = c.StatusOfCellAbove(grid)
	c.NeighbourHood.Below, _ = c.StatusOfCellBelow(grid)
	c.NeighbourHood.Left, _ = c.StatusOfCellToLeft(grid)
	c.NeighbourHood.Right, _ = c.StatusOfCellToRight(grid)
	c.NeighbourHood.AboveLeft, _ = c.StatusOfCellAboveLeft(grid)
	c.NeighbourHood.AboveRight, _ = c.StatusOfCellAboveRight(grid)
	c.NeighbourHood.BelowLeft, _ = c.StatusOfCellBelowLeft(grid)
	c.NeighbourHood.BelowRight, _ = c.StatusOfCellBelowRight(grid)
}

func (c *Cell) Update(grid *GameGrid) CellStatus {
	c.UpdateNeighbourHood(grid)

	if c.ShouldDie() {
		return Dead
	} else if c.ShouldLive() {
		return Alive
	}

	return c.Status
}

type NeighbourHood struct {
	Above      CellStatus
	Below      CellStatus
	Left       CellStatus
	Right      CellStatus
	AboveLeft  CellStatus
	AboveRight CellStatus
	BelowLeft  CellStatus
	BelowRight CellStatus
}

func (nh *NeighbourHood) GetNeighbourCount() int32 {
	count := int32(0)

	if nh.Above == Alive {
		count += 1
	}

	if nh.Below == Alive {
		count += 1
	}

	if nh.Left == Alive {
		count += 1
	}

	if nh.Right == Alive {
		count += 1
	}

	if nh.AboveLeft == Alive {
		count += 1
	}

	if nh.AboveRight == Alive {
		count += 1
	}

	if nh.BelowLeft == Alive {
		count += 1
	}

	if nh.BelowRight == Alive {
		count += 1
	}

	return count
}

func (nh *NeighbourHood) IsUnderPopulated() bool {
	return nh.GetNeighbourCount() < 2
}

func (nh *NeighbourHood) IsOverPopulated() bool {
	return nh.GetNeighbourCount() > 3
}

func (nh *NeighbourHood) IsOptimal() bool {
	return nh.GetNeighbourCount() == 3
}

func walkCells(grid *GameGrid, callback func(grid *GameGrid, cell *Cell)) {
	for rowIndex, row := range grid {
		for colIndex := range row {
			callback(grid, &grid[rowIndex][colIndex])
		}
	}
}

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Conway's Game of Life")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var grid GameGrid
	for rowIndex := int32(0); rowIndex < ROWS; rowIndex++ {
		for colIndex := int32(0); colIndex < COLUMNS; colIndex++ {
			status := Dead
			if rowIndex == 1 && colIndex == 2 ||
				rowIndex == 2 && colIndex == 2 ||
				rowIndex == 3 && colIndex == 2 {
				status = Alive
			}
			grid[rowIndex][colIndex] = Cell{Row: rowIndex, Col: colIndex, Status: status}
		}
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if rl.IsKeyPressed(rl.KeyUp) {
			newGrid := grid

			walkCells(&grid, func(grid *GameGrid, cell *Cell) {
				newGrid[cell.Row][cell.Col].Status = cell.Update(grid)
			})

			grid = newGrid
		}

		walkCells(&grid, func(grid *GameGrid, cell *Cell) {
			if cell.Status == Alive {
				rl.DrawRectangle(CELL_SIZE*cell.Col, CELL_SIZE*cell.Row, CELL_SIZE, CELL_SIZE, rl.RayWhite)
			}
			rl.DrawRectangleLines(CELL_SIZE*cell.Col, CELL_SIZE*cell.Row, CELL_SIZE, CELL_SIZE, rl.Gray)
		})

		rl.EndDrawing()
	}
}

package gameoflife

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

func (c *Cell) IsAlive() bool {
	return c.Status == Alive
}

func (c *Cell) Update(grid *Grid) CellStatus {
	c.UpdateNeighbourHood(grid)

	if c.ShouldDie() {
		return Dead
	}

	if c.ShouldLive() {
		return Alive
	}

	return c.Status
}

func (c *Cell) UpdateNeighbourHood(grid *Grid) {
	c.NeighbourHood.Above = c.StatusOfCellAbove(grid)
	c.NeighbourHood.Below = c.StatusOfCellBelow(grid)
	c.NeighbourHood.Left = c.StatusOfCellToLeft(grid)
	c.NeighbourHood.Right = c.StatusOfCellToRight(grid)
	c.NeighbourHood.AboveLeft = c.StatusOfCellAboveLeft(grid)
	c.NeighbourHood.AboveRight = c.StatusOfCellAboveRight(grid)
	c.NeighbourHood.BelowLeft = c.StatusOfCellBelowLeft(grid)
	c.NeighbourHood.BelowRight = c.StatusOfCellBelowRight(grid)
}

func (c *Cell) StatusOfCellAbove(grid *Grid) CellStatus {
	neighbour, err := grid.CellAt(c.Row-1, c.Col)
	if err != nil {
		return Dead
	}
	return neighbour.Status
}

func (c *Cell) StatusOfCellBelow(grid *Grid) CellStatus {
	neighbour, err := grid.CellAt(c.Row+1, c.Col)
	if err != nil {
		return Dead
	}
	return neighbour.Status
}

func (c *Cell) StatusOfCellToLeft(grid *Grid) CellStatus {
	neighbour, err := grid.CellAt(c.Row, c.Col-1)
	if err != nil {
		return Dead
	}
	return neighbour.Status
}

func (c *Cell) StatusOfCellToRight(grid *Grid) CellStatus {
	neighbour, err := grid.CellAt(c.Row, c.Col+1)
	if err != nil {
		return Dead
	}
	return neighbour.Status
}

func (c *Cell) StatusOfCellAboveLeft(grid *Grid) CellStatus {
	neighbour, err := grid.CellAt(c.Row-1, c.Col-1)
	if err != nil {
		return Dead
	}
	return neighbour.Status
}

func (c *Cell) StatusOfCellAboveRight(grid *Grid) CellStatus {
	neighbour, err := grid.CellAt(c.Row-1, c.Col+1)
	if err != nil {
		return Dead
	}
	return neighbour.Status
}

func (c *Cell) StatusOfCellBelowLeft(grid *Grid) CellStatus {
	neighbour, err := grid.CellAt(c.Row+1, c.Col-1)
	if err != nil {
		return Dead
	}
	return neighbour.Status
}

func (c *Cell) StatusOfCellBelowRight(grid *Grid) CellStatus {
	neighbour, err := grid.CellAt(c.Row+1, c.Col+1)
	if err != nil {
		return Dead
	}
	return neighbour.Status
}

package gameoflife

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

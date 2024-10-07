## Conway's Game of Life

Written in Go, using Raylib.

Allows user to click a cell to change it's state from _Alive_ to _Dead_.

Press the _Up Arrow_ key to show the next cycle. 

### Run code with

```
go run main.go
```

### Alter parameters in main.go

```go
const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 800
	COLS          = 16
	ROWS          = 16
	CELL_SIZE     = SCREEN_WIDTH / COLS
)
```
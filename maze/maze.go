package maze

import (
	"math/rand"

	mazelib "github.com/redwookcreek/maze/libs"
)

type Maze struct {
	Height int
	Width  int

	// Represents walls between cells
	// True if the wall exists, false otherwise
	// conversioin between cell cord to wall cord
	// are in functions below
	Walls []bool
}

// Create a new maze with walls for all cells
func NewMaze(height, width int) *Maze {
	// height and width are for cells, walls has one extra row and col
	// The first (height + 1) * width are for horizontal walls
	// and the rest height * (width + 1) are for vertical walls
	walls := (height+1)*width + height*(width+1)
	maze := Maze{height, width, make([]bool, walls)}

	for i := range maze.Walls {
		maze.Walls[i] = true
	}
	// Remove wall at entrance
	maze.Walls[0] = false
	// Remove wall at exit
	maze.Walls[maze.cell_cord_to_wall_seq(height-1, width-1, DOWN)] = false
	return &maze
}

// Enum to represent location of the wall relative to a cell
const (
	UP    = iota
	RIGHT = iota
	DOWN  = iota
	LEFT  = iota
)

// Number of horizontal walls.
func (m *Maze) horz_wall_cnt() int {
	return (m.Height + 1) * m.Width
}

// Total number of walls
func (m *Maze) WallCnt() int {
	return (m.Height+1)*m.Width + m.Height*(m.Width+1)
}

// Convert a cell coordinate to a wall number
// Decided to use a 1-dim array to represent the walls
// The first half are horizontal walls, and the second half
// are vertical walls.
func (m *Maze) cell_cord_to_wall_seq(i, j, direction int) int {
	horz_walls := m.horz_wall_cnt()
	switch direction {
	case UP:
		return i*m.Width + j
	case DOWN:
		return (i+1)*m.Width + j
	case LEFT:
		return horz_walls + i*(m.Width+1) + j
	case RIGHT:
		return horz_walls + i*(m.Width+1) + j + 1
	}
	return -1
}

func (m *Maze) WallSeqToCellCord(seq int) (int, int, int) {
	horz_walls := m.horz_wall_cnt()
	if seq < horz_walls {
		// horizontal wall
		i := seq / m.Width
		j := seq % m.Width
		if i < m.Height {
			return i, j, UP
		} else {
			return i - 1, j, DOWN
		}
	} else {
		seq = seq - horz_walls
		i := seq / (m.Width + 1)
		j := seq % (m.Width + 1)
		if j < m.Width {
			return i, j, LEFT
		} else {
			return i, j - 1, RIGHT
		}
	}
}

func (m *Maze) is_boarder_wall(i, j, dir int) bool {
	if i == 0 && dir == UP {
		return true
	}
	if i == m.Height {
		return true
	}
	if i == m.Height-1 && dir == DOWN {
		return true
	}
	if j == 0 && dir == LEFT {
		return true
	}
	if j == m.Width {
		return true
	}
	if j == m.Width-1 && dir == RIGHT {
		return true
	}
	return false
}

func get_cell_on_other_size(i, j, dir int) (int, int) {
	switch dir {
	case UP:
		return i - 1, j
	case DOWN:
		return i + 1, j
	case LEFT:
		return i, j - 1
	case RIGHT:
		return i, j + 1
	}
	// should not be called
	return -1, -1
}

// Create a maze with Iterative randomized Kruskal's algorithm
//   - randomly shuffle all the walls
//   - for each wall
//     -- if the two cells divided by the wall are not connected
//     -- remove the wall and connect the two sets
func CreateMaze(height, width int) *Maze {
	maze := NewMaze(height, width)
	uf := mazelib.CreateUnionFind(maze.Height * maze.Width)
	for _, wall_seq := range rand.Perm(maze.WallCnt()) {
		// If the cells sharing this wall belongs to different
		// set, remove the wall and join the two sets
		i1, j1, dir := maze.WallSeqToCellCord(wall_seq)
		// wall on the boarder is not shared by two cells
		if maze.is_boarder_wall(i1, j1, dir) {
			continue
		}
		i2, j2 := get_cell_on_other_size(i1, j1, dir)
		cell1 := i1*maze.Width + j1
		cell2 := i2*maze.Width + j2
		if uf.Find(cell1) == uf.Find(cell2) {
			continue
		}
		// remove the wall and join the two sets
		maze.Walls[wall_seq] = false
		uf.Union(cell1, cell2)
	}
	return maze
}

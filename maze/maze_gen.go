package maze

import (
	"math/rand"

	mazelib "github.com/redwookcreek/maze/libs"
)

// this type is used to display progress of generating a maze
// It will contain a random list of wall sequence number used
// by generation fuctions
type MazeGen struct {
	Maze

	wall_seq  []int
	next_wall int
	uf        mazelib.UnionFind
}

// Generate the maze by one step
func (maze *MazeGen) Next() {
	if maze.Finished() {
		return
	}
	for ; !maze.Finished(); maze.next_wall++ {
		// If the cells sharing this wall belongs to different
		// set, remove the wall and join the two sets
		i1, j1, dir := maze.WallSeqToCellCord(maze.wall_seq[maze.next_wall])
		// wall on the board is not shared by two cells
		if maze.is_boarder_wall(i1, j1, dir) {
			continue
		}
		i2, j2 := get_cell_on_other_size(i1, j1, dir)
		cell1 := i1*maze.Width + j1
		cell2 := i2*maze.Width + j2
		if maze.uf.Find(cell1) == maze.uf.Find(cell2) {
			continue
		}
		// remove the wall and join the two sets
		maze.Walls[maze.wall_seq[maze.next_wall]] = false
		maze.uf.Union(cell1, cell2)

		break
	}
}

func (maze *MazeGen) Finished() bool {
	return maze.next_wall >= len(maze.wall_seq)
}

func CreateMazeGen(height, width int) *MazeGen {
	maze := NewMaze(height, width)
	maze_gen := MazeGen{
		*maze,
		rand.Perm(maze.WallCnt()),
		0,
		*mazelib.CreateUnionFind(height * width),
	}
	return &maze_gen
}

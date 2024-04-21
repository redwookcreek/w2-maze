package maze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCellCordToWallSeq(t *testing.T) {

	maze := NewMaze(2, 3)

	assert.Equal(t, 0, maze.cell_cord_to_wall_seq(0, 0, UP))
	assert.Equal(t, 1, maze.cell_cord_to_wall_seq(0, 1, UP))
	assert.Equal(t, 3, maze.cell_cord_to_wall_seq(0, 0, DOWN))
	assert.Equal(t, 3, maze.cell_cord_to_wall_seq(1, 0, UP))
	assert.Equal(t, 6, maze.cell_cord_to_wall_seq(1, 0, DOWN))

	assert.Equal(t, 9, maze.cell_cord_to_wall_seq(0, 0, LEFT))
	assert.Equal(t, 10, maze.cell_cord_to_wall_seq(0, 0, RIGHT))
	assert.Equal(t, 10, maze.cell_cord_to_wall_seq(0, 1, LEFT))
	assert.Equal(t, 13, maze.cell_cord_to_wall_seq(1, 0, LEFT))
	assert.Equal(t, 15, maze.cell_cord_to_wall_seq(1, 2, LEFT))
}

// makeIS will convert any number of parameters to a []interface{}
func makeIS(v ...interface{}) []interface{} {
	return v
}

func TestWallSeqToCellCord(t *testing.T) {

	maze := NewMaze(2, 3)

	assert.Equal(t, makeIS(0, 0, UP), makeIS(maze.WallSeqToCellCord(0)))
	assert.Equal(t, makeIS(1, 0, UP), makeIS(maze.WallSeqToCellCord(3)))
	assert.Equal(t, makeIS(1, 0, DOWN), makeIS(maze.WallSeqToCellCord(6)))

	assert.Equal(t, makeIS(0, 0, LEFT), makeIS(maze.WallSeqToCellCord(9)))
	assert.Equal(t, makeIS(0, 1, LEFT), makeIS(maze.WallSeqToCellCord(10)))
	assert.Equal(t, makeIS(0, 2, LEFT), makeIS(maze.WallSeqToCellCord(11)))
	assert.Equal(t, makeIS(0, 2, RIGHT), makeIS(maze.WallSeqToCellCord(12)))
}

func TestIsBoarderWall(t *testing.T) {
	maze := NewMaze(2, 3)

	assert.True(t, maze.is_boarder_wall(0, 0, UP))
	assert.False(t, maze.is_boarder_wall(0, 0, DOWN))
	assert.True(t, maze.is_boarder_wall(0, 0, LEFT))
	assert.False(t, maze.is_boarder_wall(0, 0, RIGHT))

	assert.True(t, maze.is_boarder_wall(0, 2, RIGHT))
	assert.False(t, maze.is_boarder_wall(0, 2, LEFT))
	assert.True(t, maze.is_boarder_wall(0, 3, RIGHT))
	assert.True(t, maze.is_boarder_wall(1, 0, DOWN))
	assert.True(t, maze.is_boarder_wall(2, 0, UP))
	assert.False(t, maze.is_boarder_wall(1, 0, UP))
}

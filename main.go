package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/redwookcreek/maze/maze"
)

var (
	whiteImage = ebiten.NewImage(3, 3)

	// whiteSubImage is an internal sub image of whiteImage.
	// Use whiteSubImage at DrawTriangles instead of whiteImage in order to avoid bleeding edges.
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

	Margin = 20

	// Number of wall updates per second
	TPS = uint16(5)
)

func init() {
	whiteImage.Fill(color.White)
}

type Game struct {
	Maze     *maze.MazeGen
	tick_cnt uint16
}

func (g *Game) Update() error {
	// The update happens 60 times per second
	// We will remove 5 walls per second
	g.tick_cnt = (g.tick_cnt + 1) % 60
	if !g.Maze.Finished() && g.tick_cnt%(60/TPS) == 0 {
		g.Maze.Next()
	}
	return nil
}

func get_wall_cord_from_cell(cell_w, cell_h, i, j, dir int) (start_x, start_y, end_x, end_y float32) {
	var x1, x2, y1, y2 float32
	switch dir {
	case maze.UP:
		y1 = float32(i * cell_h)
		x1 = float32(j * cell_w)
		y2 = y1
		x2 = x1 + float32(cell_w)
	case maze.DOWN:
		y1 = float32((i + 1) * cell_h)
		x1 = float32(j * cell_w)
		y2 = y1
		x2 = x1 + float32(cell_w)
	case maze.LEFT:
		y1 = float32(i * cell_h)
		x1 = float32(j * cell_w)
		y2 = y1 + float32(cell_h)
		x2 = x1
	case maze.RIGHT:
		y1 = float32(i * cell_h)
		x1 = float32((j + 1) * cell_w)
		y2 = y1 + float32(cell_h)
		x2 = x1
	}
	return x1, y1, x2, y2
}

// Draw all the walls
func (g *Game) Draw(screen *ebiten.Image) {
	width := screen.Bounds().Dx() - Margin // leave some margin
	height := screen.Bounds().Dy() - Margin
	cell_width := width / g.Maze.Width
	cell_height := height / g.Maze.Height
	var path vector.Path
	for wall_seq, has_wall := range g.Maze.Walls {
		if has_wall {
			i, j, dir := g.Maze.WallSeqToCellCord(wall_seq)
			x1, y1, x2, y2 := get_wall_cord_from_cell(cell_width, cell_height, i, j, dir)
			path.MoveTo(x1, y1)
			path.LineTo(x2, y2)
		}
	}

	var vs []ebiten.Vertex
	var is []uint16
	op := &vector.StrokeOptions{}
	op.Width = 1
	op.LineJoin = vector.LineJoinRound
	vs, is = path.AppendVerticesAndIndicesForStroke(nil, nil, op)

	for i := range vs {
		vs[i].DstX = vs[i].DstX + float32(Margin/2)
		vs[i].DstY = vs[i].DstY + float32(Margin/2)
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 1
		vs[i].ColorG = 1
		vs[i].ColorB = 0
		vs[i].ColorA = 1
	}

	draw_op := &ebiten.DrawTrianglesOptions{}
	draw_op.AntiAlias = true
	draw_op.FillRule = ebiten.NonZero
	screen.DrawTriangles(vs, is, whiteSubImage, draw_op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s := min(outsideWidth, outsideHeight)
	return s - 10, s - 10
}

func main() {
	m := maze.CreateMazeGen(10, 10)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Maze")
	if err := ebiten.RunGame(&Game{m, 0}); err != nil {
		log.Fatal(err)
	}
}

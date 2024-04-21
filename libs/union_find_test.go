package mazelib_test

import (
	"fmt"
	"testing"

	mazelib "github.com/redwookcreek/maze/libs"
	"github.com/stretchr/testify/assert"
)

func TestUnionFind(t *testing.T) {
	uf := mazelib.CreateUnionFind(10)

	assert.Equal(t, 0, uf.Find(0))
	assert.Equal(t, 1, uf.Find(1))

	uf.Union(0, 1)
	assert.Equal(t, 1, uf.Find(0))
	assert.Equal(t, 1, uf.Find(1))
	assert.Equal(t, 0, uf.Weight[0])
	assert.Equal(t, 2, uf.Weight[1])

	uf.Union(0, 2)
	assert.Equal(t, 1, uf.Find(2))
	assert.Equal(t, 0, uf.Weight[2])
	assert.Equal(t, 3, uf.Weight[1])

	uf.Union(3, 4)
	assert.Equal(t, 4, uf.Find(3))
	assert.Equal(t, 4, uf.Find(4))
}

func TestSet(t *testing.T) {
	uf := mazelib.CreateUnionFind(20)
	for i := 0; i < 10; i++ {
		uf.Union(0, i)
	}
	for i := 0; i < 10; i++ {
		assert.Equal(
			t, 1, uf.Find(i),
			fmt.Sprintf("Union of %d is %d", i, uf.Find(i)))
	}

	for i := 10; i < 20; i++ {
		uf.Union(10, i)
	}
	for i := 10; i < 20; i++ {
		assert.Equal(
			t, 11, uf.Find(i),
			fmt.Sprintf("Union of %d is %d", i, uf.Find(i)))
	}

	for i := 0; i < 20; i++ {
		for j := i + 1; j < 20; j++ {
			if i < 10 && j < 10 || i >= 10 && j >= 10 {
				assert.Equal(t, uf.Find(i), uf.Find(j))
			} else {
				assert.NotEqual(t, uf.Find(i), uf.Find(j))
			}
		}
	}

	for i := 0; i < 20; i++ {
		if i == 1 || i == 11 {
			assert.Equal(t, 10, uf.Weight[i])
		} else {
			assert.Equal(t, 0, uf.Weight[i])
		}
	}
}

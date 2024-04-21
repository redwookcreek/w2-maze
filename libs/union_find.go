package mazelib

type UnionFind struct {
	Parent []int
	Weight []int
}

func CreateUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	weight := make([]int, n)
	for i := range parent {
		parent[i] = i
		weight[i] = 1
	}
	return &UnionFind{parent, weight}
}

// Return the parent of node n
func (uf *UnionFind) Find(n int) int {
	for uf.Parent[n] != n {
		n = uf.Parent[n]
	}
	return n
}

// Union partition of the two nodes
// return the parent after union
func (uf *UnionFind) Union(s, t int) int {
	sp := uf.Find(s)
	tp := uf.Find(t)
	if sp == tp {
		return sp
	}
	mp := tp
	mt := s
	if uf.Weight[sp] > uf.Weight[tp] {
		mp = sp
		mt = t
	}
	for uf.Parent[mt] != mt {
		next_parent := uf.Parent[mt]
		uf.Parent[mt] = mp
		mt = next_parent
	}
	uf.Parent[mt] = mp
	uf.Weight[mp] += uf.Weight[mt]
	uf.Weight[mt] = 0
	return mp
}

package unionfind

type UnionFind struct {
	Parent        []int
	Size          []int
	NumComponents int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	numComponents := n

	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	return &UnionFind{Parent: parent, Size: size, NumComponents: numComponents}
}

func (uf *UnionFind) Find(i int) int {
	if uf.Parent[i] == i {
		return i
	}

	uf.Parent[i] = uf.Find(uf.Parent[i]) // path compression
	return uf.Parent[i]
}

func (uf *UnionFind) Union(i int, j int) bool {
	root_i := uf.Find(i)
	root_j := uf.Find(j)

	if root_i == root_j {
		return false
	}

	if uf.Size[root_i] >= uf.Size[root_j] {
		uf.Parent[root_j] = root_i
		uf.Size[root_i] += uf.Size[root_j]
	} else {
		uf.Parent[root_i] = root_j
		uf.Size[root_j] += uf.Size[root_i]
	}

	uf.NumComponents--

	return true
}

func (uf *UnionFind) ComponentSize(i int) int {
	return uf.Size[uf.Find(i)]
}

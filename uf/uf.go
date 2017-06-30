package uf

// UF is union find
type UF struct {
	id []int
	sz []int
}

// NewUF returns a new UF
func NewUF(n int) *UF {
	uf := new(UF)
	uf.id = make([]int, n)
	uf.sz = make([]int, n)

	for i := range uf.id {
		uf.id[i] = i
		uf.sz[i] = 1
	}
	return uf
}

func (uf UF) root(i int) int {
	for i != uf.id[i] {
		uf.id[i] = uf.id[uf.id[i]]
		i = uf.id[i]
	}
	return i
}

// Union add connection between p and q
func (uf UF) Union(p, q int) {
	i, j := uf.root(p), uf.root(q)
	if uf.sz[i] < uf.sz[j] {
		uf.id[i] = j
		uf.sz[j] += uf.sz[i]
	} else {
		uf.id[j] = i
		uf.sz[i] += uf.sz[j]
	}
	uf.id[i] = j
}

// Connected whether p and q in the same component
func (uf UF) Connected(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

// Find component identifier for p (0 to N – 1)
func (uf UF) Find(p int) int {
	return uf.root(p)
}

// Count return number of componets
func (uf UF) Count() int {
	return len(uf.id)
}

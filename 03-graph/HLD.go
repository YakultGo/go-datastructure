type HLD struct {
	n      int
	parent []int
	dep    []int
	sz     []int
	son    []int
	top    []int
	in     []int
	out    []int
	rnk    []int
	adj    [][]int
	cnt    int
}

func NewHLD(n int) *HLD {
	hld := &HLD{n: n}
	hld.init(n)
	return hld
}

func (hld *HLD) init(n int) {
	hld.n = n
	hld.parent = make([]int, n+1)
	hld.dep = make([]int, n+1)
	hld.sz = make([]int, n+1)
	hld.son = make([]int, n+1)
	hld.top = make([]int, n+1)
	hld.in = make([]int, n+1)
	hld.out = make([]int, n+1)
	hld.rnk = make([]int, n+1)
	hld.cnt = 0
	hld.adj = make([][]int, n+1)
	for i := range hld.adj {
		hld.adj[i] = []int{}
	}
}

func (hld *HLD) addEdge(u, v int) {
	hld.adj[u] = append(hld.adj[u], v)
	hld.adj[v] = append(hld.adj[v], u)
}

func (hld *HLD) work(root int) {
	hld.dep[0] = -1
	hld.dfs1(root, 0)
	hld.dfs2(root, root)
}

func (hld *HLD) dfs1(x, fa int) {
	hld.parent[x] = fa
	hld.dep[x] = hld.dep[fa] + 1
	hld.sz[x] = 1
	for _, y := range hld.adj[x] {
		if y == fa {
			continue
		}
		hld.dfs1(y, x)
		hld.sz[x] += hld.sz[y]
		if hld.sz[y] > hld.sz[hld.son[x]] {
			hld.son[x] = y
		}
	}
}

func (hld *HLD) dfs2(x, r int) {
	hld.top[x] = r
	hld.cnt++
	hld.in[x] = hld.cnt
	hld.rnk[hld.cnt] = x
	if hld.son[x] == 0 {
		hld.out[x] = hld.cnt
		return
	}
	hld.dfs2(hld.son[x], r)
	for _, y := range hld.adj[x] {
		if y == hld.parent[x] || y == hld.son[x] {
			continue
		}
		hld.dfs2(y, y)
	}
	hld.out[x] = hld.cnt
}

func (hld *HLD) lca(u, v int) int {
	for hld.top[u] != hld.top[v] {
		if hld.dep[hld.top[u]] < hld.dep[hld.top[v]] {
			u, v = v, u
		}
		u = hld.parent[hld.top[u]]
	}
	if hld.dep[u] < hld.dep[v] {
		return u
	}
	return v
}

func (hld *HLD) dis(u, v int) int {
	return hld.dep[u] + hld.dep[v] - 2*hld.dep[hld.lca(u, v)]
}

func (hld *HLD) rootedLCA(a, b, c int) int {
	return hld.lca(a, b) ^ hld.lca(b, c) ^ hld.lca(c, a)
}

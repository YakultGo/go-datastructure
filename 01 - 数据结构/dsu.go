package datastructure

type Dsu struct {
    Fa     []int
    Size   []int
    Groups int
}

func NewDsu(n int) *Dsu {
    fa := make([]int, n)
    size := make([]int, n)
    for i := 0; i < n; i++ {
        fa[i] = i
        size[i] = 1
    }
    return &Dsu{Fa: fa, Size: size, Groups: n}
}
func (u *Dsu) Find(x int) int {
    if u.Fa[x] != x {
        u.Fa[x] = u.Find(u.Fa[x])
    }
    return u.Fa[x]
}

func (u *Dsu) Union(x, y int) {
    x = u.Find(x)
    y = u.Find(y)
    if x == y {
        return
    }
    if u.Size[x] < u.Size[y] {
        x, y = y, x
    }
    u.Fa[y] = x
    u.Size[x] += u.Size[y]
    u.Groups--
}

func (u *Dsu) Same(x, y int) bool {
    return u.Find(x) == u.Find(y)
}

package datastructure

type FenwickTree struct {
    tr []int
}

func NewFenwickTree(n int) *FenwickTree {
    return &FenwickTree{tr: make([]int, n+10)}
}
func lowbit(x int) int {
    return x & -x
}
func (ft *FenwickTree) add(x, delta int) {
    for x < len(ft.tr) {
        ft.tr[x] += delta
        x += lowbit(x)
    }
}

func (ft *FenwickTree) query(x int) int {
    var res int
    for x > 0 {
        res += ft.tr[x]
        x -= lowbit(x)
    }
    return res
}

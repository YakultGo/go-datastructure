type LazySeg []struct {
	l, r int
	val  int
	tag  int
}

func (t LazySeg) merge(a, b int) int {
	return max(a, b)
}
func (t LazySeg) info(o int, v int) {
	t[o].val += v
	t[o].tag += v
}
func (t LazySeg) maintain(o int) {
	t[o].val = t.merge(t[o<<1].val, t[o<<1|1].val)
}
func (t LazySeg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].val = int(-1e18)
	if l == r {
		t[o].val = a[l]
		return
	}
	m := (l + r) >> 1
	if l <= m {
		t.build(a, o<<1, l, m)
	}
	if m < r {
		t.build(a, o<<1|1, m+1, r)
	}
	t.maintain(o)
}
func (t LazySeg) pushDown(o int) {
	if t[o].tag != 0 {
		t.info(o<<1, t[o].tag)
		t.info(o<<1|1, t[o].tag)
		t[o].tag = 0
	}
}
func (t LazySeg) update(o, l, r int, delta int) {
	if l <= t[o].l && t[o].r <= r {
		t.info(o, delta)
		return
	}
	t.pushDown(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, delta)
	}
	if m < r {
		t.update(o<<1|1, l, r, delta)
	}
	t.maintain(o)
}

func (t LazySeg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	t.pushDown(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	lRes := t.query(o<<1, l, r)
	rRes := t.query(o<<1|1, l, r)
	return t.merge(lRes, rRes)
}
func newLazySegmentTree(a []int) LazySeg {
	n := len(a)
	t := make(LazySeg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}
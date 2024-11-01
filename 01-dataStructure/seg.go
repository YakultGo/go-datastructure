type Seg []struct {
	l, r int
	val  int
}

func (t Seg) info(o int, v int) {
	t[o].val += v
}
func (t Seg) maintain(o int) {
	t[o].val = t[o<<1].val + t[o<<1|1].val
}
func (t Seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
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

func (t Seg) update(o, pos int, delta int) {
	if t[o].l == t[o].r {
		t.info(o, delta)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if pos <= m {
		t.update(o<<1, pos, delta)
	}
	if m < pos {
		t.update(o<<1|1, pos, delta)
	}
	t.maintain(o)
}

func newSegmentTree(a []int) Seg {
	n := len(a)
	t := make(Seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}
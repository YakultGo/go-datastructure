type RMQ struct {
	n     int
	a     []int
	f     [][]int
	_func func(int, int) int
}

func (rmq *RMQ) init() {
	n := rmq.n
	rmq.f = make([][]int, n)
	len := 63 - bits.LeadingZeros(uint(n)) + 1
	for i := 0; i < n; i++ {
		rmq.f[i] = make([]int, len)
		rmq.f[i][0] = rmq.a[i]
	}
	for j := 1; (1 << j) <= n; j++ {
		for i := 0; i+(1<<j)-1 < n; i++ {
			rmq.f[i][j] = rmq._func(rmq.f[i][j-1], rmq.f[i+(1<<(j-1))][j-1])
		}
	}
}

func (rmq *RMQ) Query(l, r int) int {
	k := 63 - bits.LeadingZeros(uint(r-l+1))
	return rmq._func(rmq.f[l][k], rmq.f[r-(1<<k)+1][k])
}

func NewRMQ(a []int, cmp func(int, int) int) *RMQ {
	if cmp == nil {
		cmp = func(a, b int) int {
			return max(a, b)
		}
	}
	rmq := &RMQ{
		a:     a,
		n:     len(a),
		_func: cmp,
	}
	rmq.init()
	return rmq
}
type node struct {
    value int
}
type hp []*node

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].value < h[j].value }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(*node)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v *node)      { heap.Push(h, v) }
func (h *hp) pop() *node        { return heap.Pop(h).(*node) }
func (h *hp) top() *node        { return (*h)[0] }
func (h *hp) empty() bool       { return len(*h) == 0 }
func (h *hp) size() int         { return len(*h) }
func NewPriorityQueue() *hp     { h := hp(make([]*node, 0)); heap.Init(&h); return &h }

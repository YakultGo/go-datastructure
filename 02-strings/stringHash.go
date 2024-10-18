var Mod = []int64{1610612741, 0, 805306457, 402653189, 201326611, 100663319, 50331653, 1000000007, 1000000009}

type Hash struct {
    n  int
    M1 []int64
    M2 []int64
    h1 []int64
    h2 []int64
    p1 int64
    p2 int64
}

func NewHash(s string) *Hash {
    h := &Hash{
        p1: 131,
        p2: 13331,
    }
    h.Init(s)
    return h
}

func (h *Hash) Init(s string) {
    if Mod[1] == 0 {
        h.change()
    }
    h.n = len(s)
    h.work(1, h.n, s)
}

func (h *Hash) change() {
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))
    for Mod[1] == 0 || Mod[2] == 0 {
        rng.Shuffle(len(Mod), func(i, j int) {
            Mod[i], Mod[j] = Mod[j], Mod[i]
        })
    }
}

func (h *Hash) work(st, length int, s string) {
    h.M1 = make([]int64, length+1)
    h.M2 = make([]int64, length+1)
    h.h1 = make([]int64, length+1)
    h.h2 = make([]int64, length+1)
    h.M1[0] = 1
    h.M2[0] = 1
    for i := st; i <= length; i++ {
        h.M1[i] = (h.M1[i-1] * h.p1) % Mod[1]
        h.h1[i] = (h.h1[i-1]*h.p1 + int64(s[i-st])) % Mod[1]
        h.M2[i] = (h.M2[i-1] * h.p2) % Mod[2]
        h.h2[i] = (h.h2[i-1]*h.p2 + int64(s[i-st])) % Mod[2]
    }
}

func (h *Hash) Get(l, r int) (int64, int64) {
    l++
    r++
    t1 := ((h.h1[r]-h.h1[l-1]*h.M1[r-l+1])%Mod[1] + Mod[1]) % Mod[1]
    t2 := ((h.h2[r]-h.h2[l-1]*h.M2[r-l+1])%Mod[2] + Mod[2]) % Mod[2]
    return t1, t2
}

func (h *Hash) Mul(l, r [2]int64, length int64) (int64, int64) {
    l[0] = (l[0] * h.M1[length]) % Mod[1]
    l[1] = (l[1] * h.M2[length]) % Mod[2]
    it1 := (l[0] + r[0]) % Mod[1]
    it2 := (l[1] + r[1]) % Mod[2]
    return it1, it2
}

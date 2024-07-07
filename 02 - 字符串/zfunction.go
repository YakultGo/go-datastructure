func zFunction(s string) []int {
    n := len(s)
    z := make([]int, n)
    l, r := 0, 0
    for i := 1; i < n; i++ {
        if i <= r && z[i-l] < r-i+1 {
            z[i] = z[i-l]
        } else {
            z[i] = max(0, r-i+1)
            for i+z[i] < n && s[z[i]] == s[i+z[i]] {
                z[i]++
            }
        }
        if i+z[i]-1 > r {
            l, r = i, i+z[i]-1
        }
    }
    return z
}
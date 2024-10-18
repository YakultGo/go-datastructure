func prefixFunction(s string) []int {
    n := len(s)
    pi := make([]int, n)
    for i := 1; i < n; i++ {
        j := pi[i-1]
        for j > 0 && s[i] != s[j] {
            j = pi[j-1]
        }
        if s[i] == s[j] {
            j++
        }
        pi[i] = j
    }
    return pi
}

func kmp(text, pattern string) []int {
    pi := prefixFunction(pattern)
    var res []int
    c, m := 0, len(pattern)
    for i := 0; i < len(text); i++ {
        v := text[i]
        for c > 0 && pattern[c] != v {
            c = pi[c-1]
        }
        if pattern[c] == v {
            c++
        }
        if c == m {
            res = append(res, i-m+1)
            c = pi[c-1]
        }
    }
    return res
}
func nex(s string) []int {
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
	pi := nex(pattern)
	var res []int
	j, m := 0, len(pattern)
	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] == pattern[j] {
			j = pi[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			res = append(res, i-m+1)
			j = pi[j-1]
		}
	}
	return res
}
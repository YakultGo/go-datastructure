// 全排列函数
func permute(n int) [][]int {
	var result [][]int
	var dfs func([]int)
	vis := make([]bool, n)
	dfs = func(path []int) {
		if len(path) == n {
			tmp := make([]int, n)
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if vis[i] {
				continue
			}
			vis[i] = true
			path = append(path, i)
			dfs(path)
			path = path[:len(path)-1]
			vis[i] = false
		}
	}
	dfs([]int{})
	return result
}

// gcd 最大公约数
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// ksm 快速幂
func ksm(a, b, p int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a % p
		}
		a = a * a % p
		b >>= 1
	}
	return res
}
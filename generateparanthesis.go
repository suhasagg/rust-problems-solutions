func generateParenthesis(n int) []string {
   var res []string
	var b []byte = make([]byte, 2*n)
	var dfs func(int, int, int)
	dfs = func(l, r, idx int) {
		if l == n && r == n {
			res = append(res, string(b))
			return
		}
		if l < n {
			b[idx] = '('
			dfs(l+1, r, idx+1)
		}
		if r < l {
			b[idx] = ')'
			dfs(l, r+1, idx+1)
		}
	}
	dfs(0, 0, 0)
	return res
}
